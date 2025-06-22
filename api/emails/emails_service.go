package emails

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"portfolio/api/config"
	"strings"
	"time"

	openai "github.com/sashabaranov/go-openai"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type EmailsService struct {
	mongoDB *mongo.Database
}

func NewEmailsService(mongoDB *mongo.Database) *EmailsService {
	return &EmailsService{
		mongoDB: mongoDB,
	}
}

func (ctx *EmailsService) SendPortfolioMessage(data SendPortfolioMessage) (string, error) {
	from := mail.NewEmail("Portfolio Message", "no.reply@takedi.com")
	subject := "New message from portfolio 🎉"
	to := mail.NewEmail("Vinicius Takedi", "viniciustakedi7@gmail.com")

	plainTextContent := getPortfolioMessagePlainText(data)
	htmlContent := getPortfolioMessageHTML(data)

	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(config.GetEnv("SENDGRID_API_KEY"))

	response, err := client.Send(message)
	if err != nil {
		return "", err
	}

	if response.StatusCode != 202 {
		return "", errors.New(response.Body)
	}

	return "Email sent successfully!", nil
}

func (svc *EmailsService) GetNewsletterScheduleTime(newsletterId string) (string, error) {
	objectId, err := primitive.ObjectIDFromHex(newsletterId)
	if err != nil {
		return "", fmt.Errorf("invalid newsletterId: %w", err)
	}

	ctx := context.Background()

	newsletterColl := svc.mongoDB.Collection("newsletters")
	newsletterDb, err := newsletterColl.Find(ctx, bson.M{
		"_id": objectId,
	})
	if err != nil {
		return "", fmt.Errorf("finding newsletter: %w", err)
	}
	defer newsletterDb.Close(ctx)

	var newsletter struct {
		ID       primitive.ObjectID `bson:"_id"`
		Name     string             `bson:"name"`
		Schedule string             `bson:"schedule"`
	}

	if !newsletterDb.Next(ctx) {
		return "", fmt.Errorf("newsletter not found")
	}

	if err := newsletterDb.Decode(&newsletter); err != nil {
		return "", fmt.Errorf("decoding newsletter: %w", err)
	}

	return newsletter.Schedule, nil
}

func (svc *EmailsService) SendDailyWordNewsletter() error {
	// 1 - Prepare context and IDs & Get word of the day
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	newsletterTypeID, err := primitive.ObjectIDFromHex("684cd13895298f80e21813a9")
	if err != nil {
		return fmt.Errorf("invalid newsletterTypeId: %w", err)
	}

	wordColl := svc.mongoDB.Collection("dailywordnewsletterwords")
	filter := bson.M{"used": false}
	update := bson.M{"$set": bson.M{"used": true}}
	// return the *updated* document
	opts := options.FindOneAndUpdate().
		SetReturnDocument(options.After)

	wordRes := wordColl.FindOneAndUpdate(ctx, filter, update, opts)
	if err := wordRes.Err(); err != nil {
		if err == mongo.ErrNoDocuments {
			return fmt.Errorf("no unused words left")
		}
		return fmt.Errorf("findOneAndUpdate error: %w", err)
	}

	var word WordDB
	if err := wordRes.Decode(&word); err != nil {
		return fmt.Errorf("decoding word: %w", err)
	}

	// 2 - Load API keys
	openAIKey := config.GetEnv("OPENAI_API_KEY")
	if openAIKey == "" {
		return errors.New("OPENAI_API_KEY not set")
	}
	sendGridKey := config.GetEnv("SENDGRID_API_KEY")
	if sendGridKey == "" {
		return errors.New("SENDGRID_API_KEY not set")
	}

	// 3 - Build OpenAI request
	client := openai.NewClient(openAIKey)
	req := openai.ChatCompletionRequest{
		Model: openai.GPT4,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: `You are a dictionary assistant. Respond with JSON only, containing these fields: word (string), definition (string), usageTip (string), funFact (string), examples (array of strings), synonyms (array of strings), antonyms (array of strings).`,
			},
			{
				Role:    openai.ChatMessageRoleUser,
				Content: fmt.Sprintf("Provide detailed information for the word %s", word.Name),
			},
		},
		Temperature:      0.9,
		TopP:             0.9,
		PresencePenalty:  0.6,
		FrequencyPenalty: 0.6,
	}

	// 4 -  Call with retries on 5xx
	var chatResp openai.ChatCompletionResponse
	for attempt := 1; attempt <= 3; attempt++ {
		chatResp, err = client.CreateChatCompletion(ctx, req)
		if err == nil {
			break
		}
		var apiErr *openai.APIError
		if errors.As(err, &apiErr) && apiErr.HTTPStatusCode >= 500 && apiErr.HTTPStatusCode < 600 {
			time.Sleep(time.Duration(attempt) * time.Second)
			continue
		}
		return fmt.Errorf("OpenAI request error: %w", err)
	}
	if err != nil {
		return fmt.Errorf("OpenAI request failed after retries: %w", err)
	}

	// 5 -  Parse JSON safely
	content := chatResp.Choices[0].Message.Content
	var info WordInfo
	if err := json.Unmarshal([]byte(content), &info); err != nil {
		return fmt.Errorf("invalid JSON in OpenAI response: %w\nraw: %s", err, content)
	}

	// Capitalize the word
	info.Word = strings.ToUpper(info.Word[:1]) + info.Word[1:]

	// 6 - Insert into MongoDB
	coll := svc.mongoDB.Collection("dailywordnewsletter")
	now := time.Now()
	doc := bson.M{
		"newsletterTypeId": newsletterTypeID,
		"word":             info.Word,
		"definition":       info.Definition,
		"usageTip":         info.UsageTip,
		"funFact":          info.FunFact,
		"examples":         info.Examples,
		"synonyms":         info.Synonyms,
		"antonyms":         info.Antonyms,
		"createdAt":        now,
		"sentAt":           nil,
	}
	res, err := coll.InsertOne(ctx, doc)
	if err != nil {
		return fmt.Errorf("insert newsletter: %w", err)
	}
	newsletterID := res.InsertedID.(primitive.ObjectID)

	// 7 - Load recipients
	recColl := svc.mongoDB.Collection("recipients")
	cursor, err := recColl.Find(ctx, bson.M{
		"newsletterTypeId": newsletterTypeID,
		"subscribed":       true,
	})
	if err != nil {
		return fmt.Errorf("finding recipients: %w", err)
	}
	defer cursor.Close(ctx)

	var recipients []struct {
		ID    primitive.ObjectID `bson:"_id"`
		Name  string             `bson:"name"`
		Email string             `bson:"email"`
	}
	if err := cursor.All(ctx, &recipients); err != nil {
		return fmt.Errorf("decoding recipients: %w", err)
	}

	if len(recipients) == 0 {
		return fmt.Errorf("no recipients was found")
	}

	// 8 - Build and send via SendGrid
	sgClient := sendgrid.NewSendClient(sendGridKey)
	from := mail.NewEmail("Cacau says whoooof!", "no.reply@takedi.com")
	subject := fmt.Sprintf("Learn with Cacau — %s", info.Word)

	m := mail.NewV3Mail()
	m.SetFrom(from)
	m.Subject = subject
	m.AddContent(mail.NewContent("text/plain", getDailyWordNewsletterPlainText(info)))
	m.AddContent(mail.NewContent("text/html", getDailyWordNewsletterHTML(info)))

	for _, r := range recipients {
		p := mail.NewPersonalization()
		p.AddTos(mail.NewEmail(r.Name, r.Email))
		m.AddPersonalizations(p)
	}

	respSG, err := sgClient.Send(m)
	if err != nil {
		return fmt.Errorf("SendGrid error: %w", err)
	}
	if respSG.StatusCode >= 400 {
		return fmt.Errorf("SendGrid API error: %s", respSG.Body)
	}

	// 9 - Mark as sent
	_, err = coll.UpdateByID(ctx, newsletterID, bson.M{"$set": bson.M{"sentAt": time.Now()}})
	if err != nil {
		return fmt.Errorf("update newsletter: %w", err)
	}

	return nil
}

func (svc *EmailsService) SendDailyPhrasalVerbNewsletter() error {
	// 1 - Prepare context and IDs & Get word of the day
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	newsletterTypeID, err := primitive.ObjectIDFromHex("684cd13895298f80e21813a9")
	if err != nil {
		return fmt.Errorf("invalid newsletterTypeId: %w", err)
	}

	phrasalVerbColl := svc.mongoDB.Collection("dailywordnewsletterphrasalverbs")
	filter := bson.M{"used": false}
	update := bson.M{"$set": bson.M{"used": true}}
	// return the *updated* document
	opts := options.FindOneAndUpdate().
		SetReturnDocument(options.After)

	phrasalVerbRes := phrasalVerbColl.FindOneAndUpdate(ctx, filter, update, opts)
	if err := phrasalVerbRes.Err(); err != nil {
		if err == mongo.ErrNoDocuments {
			return fmt.Errorf("no unused words left")
		}
		return fmt.Errorf("findOneAndUpdate error: %w", err)
	}

	var phrasalVerb WordDB
	if err := phrasalVerbRes.Decode(&phrasalVerb); err != nil {
		return fmt.Errorf("decoding phrasalverb: %w", err)
	}

	// 2 - Load API keys
	openAIKey := config.GetEnv("OPENAI_API_KEY")
	if openAIKey == "" {
		return errors.New("OPENAI_API_KEY not set")
	}
	sendGridKey := config.GetEnv("SENDGRID_API_KEY")
	if sendGridKey == "" {
		return errors.New("SENDGRID_API_KEY not set")
	}

	// 3 - Build OpenAI request
	client := openai.NewClient(openAIKey)
	req := openai.ChatCompletionRequest{
		Model: openai.GPT4,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: `You are a dictionary assistant. Respond with JSON only, containing these fields: word (string), definition (string), usageTip (string), funFact (string), examples (array of strings), synonyms (array of strings), antonyms (array of strings).`,
			},
			{
				Role:    openai.ChatMessageRoleUser,
				Content: fmt.Sprintf("Provide detailed information for the phrasal verb %s", phrasalVerb.Name),
			},
		},
		Temperature:      0.9,
		TopP:             0.9,
		PresencePenalty:  0.6,
		FrequencyPenalty: 0.6,
	}

	// 4 -  Call with retries on 5xx
	var chatResp openai.ChatCompletionResponse
	for attempt := 1; attempt <= 3; attempt++ {
		chatResp, err = client.CreateChatCompletion(ctx, req)
		if err == nil {
			break
		}
		var apiErr *openai.APIError
		if errors.As(err, &apiErr) && apiErr.HTTPStatusCode >= 500 && apiErr.HTTPStatusCode < 600 {
			time.Sleep(time.Duration(attempt) * time.Second)
			continue
		}
		return fmt.Errorf("OpenAI request error: %w", err)
	}
	if err != nil {
		return fmt.Errorf("OpenAI request failed after retries: %w", err)
	}

	// 5 -  Parse JSON safely
	content := chatResp.Choices[0].Message.Content
	var info WordInfo
	if err := json.Unmarshal([]byte(content), &info); err != nil {
		return fmt.Errorf("invalid JSON in OpenAI response: %w\nraw: %s", err, content)
	}

	// Capitalize the word
	info.Word = strings.ToUpper(info.Word[:1]) + info.Word[1:]

	// 6 - Insert into MongoDB
	coll := svc.mongoDB.Collection("dailywordnewsletter")
	now := time.Now()
	doc := bson.M{
		"newsletterTypeId": newsletterTypeID,
		"word":             info.Word,
		"definition":       info.Definition,
		"usageTip":         info.UsageTip,
		"funFact":          info.FunFact,
		"examples":         info.Examples,
		"synonyms":         info.Synonyms,
		"antonyms":         info.Antonyms,
		"createdAt":        now,
		"sentAt":           nil,
	}
	res, err := coll.InsertOne(ctx, doc)
	if err != nil {
		return fmt.Errorf("insert newsletter: %w", err)
	}
	newsletterID := res.InsertedID.(primitive.ObjectID)

	// 7 - Load recipients
	recColl := svc.mongoDB.Collection("recipients")
	cursor, err := recColl.Find(ctx, bson.M{
		"newsletterTypeId": newsletterTypeID,
		"subscribed":       true,
	})
	if err != nil {
		return fmt.Errorf("finding recipients: %w", err)
	}
	defer cursor.Close(ctx)

	var recipients []struct {
		ID    primitive.ObjectID `bson:"_id"`
		Name  string             `bson:"name"`
		Email string             `bson:"email"`
	}
	if err := cursor.All(ctx, &recipients); err != nil {
		return fmt.Errorf("decoding recipients: %w", err)
	}

	if len(recipients) == 0 {
		return fmt.Errorf("no recipients was found")
	}

	// 8 - Build and send via SendGrid
	sgClient := sendgrid.NewSendClient(sendGridKey)
	from := mail.NewEmail("Cacau says whoooof!", "no.reply@takedi.com")
	subject := fmt.Sprintf("Learn with Cacau — %s", info.Word)

	m := mail.NewV3Mail()
	m.SetFrom(from)
	m.Subject = subject
	m.AddContent(mail.NewContent("text/plain", getDailyPhrasalVerbNewsletterPlainText(info)))
	m.AddContent(mail.NewContent("text/html", getDailyPhrasalVerbNewsletterHTML(info)))

	for _, r := range recipients {
		p := mail.NewPersonalization()
		p.AddTos(mail.NewEmail(r.Name, r.Email))
		m.AddPersonalizations(p)
	}

	respSG, err := sgClient.Send(m)
	if err != nil {
		return fmt.Errorf("SendGrid error: %w", err)
	}
	if respSG.StatusCode >= 400 {
		return fmt.Errorf("SendGrid API error: %s", respSG.Body)
	}

	// 9 - Mark as sent
	_, err = coll.UpdateByID(ctx, newsletterID, bson.M{"$set": bson.M{"sentAt": time.Now()}})
	if err != nil {
		return fmt.Errorf("update newsletter: %w", err)
	}

	return nil
}
