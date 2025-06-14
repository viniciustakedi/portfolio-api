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

func (ctx *EmailsService) SendDailyWordNewsletter() error {
	ctxBg := context.Background()
	newsletterTypeId, err := primitive.ObjectIDFromHex("684cd13895298f80e21813a9")
	if err != nil {
		return fmt.Errorf("invalid newsletterTypeId: %s", err.Error())
	}

	openAIApiKey := config.GetEnv("OPENAI_API_KEY")
	if openAIApiKey == "" {
		return errors.New("OPENAI_API_KEY not set")
	}

	client := openai.NewClient(openAIApiKey)

	system := openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleSystem,
		Content: "You are a dictionary assistant. Respond with JSON only, containing these fields: word (string), definition (string), usageTip (string), funFact (string), examples (array of strings), synonyms (array of strings), antonyms (array of strings).",
	}

	user := openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: "Provide detailed information for some random word that you want to give to me! Could be any kind of word but must be in English.",
	}

	resp, err := client.CreateChatCompletion(ctxBg, openai.ChatCompletionRequest{
		Model:            openai.GPT4,
		Messages:         []openai.ChatCompletionMessage{system, user},
		Temperature:      0.9,
		TopP:             0.9,
		PresencePenalty:  0.6,
		FrequencyPenalty: 0.6,
	})
	if err != nil {
		return fmt.Errorf("OpenAI request error: %s", err.Error())
	}

	var info WordInfo
	if err := json.Unmarshal([]byte(resp.Choices[0].Message.Content), &info); err != nil {
		return fmt.Errorf("JSON unmarshal error: %w\nresponse was: %s", err, resp.Choices[0].Message.Content)
	}

	info.Word = strings.ToUpper(info.Word[:1]) + info.Word[1:]

	dailyWordNewsletterCollection := ctx.mongoDB.Collection("dailywordnewsletter")

	now := time.Now()
	doc := bson.M{
		"newsletterTypeId": newsletterTypeId,
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
	insertRes, err := dailyWordNewsletterCollection.InsertOne(ctxBg, doc)
	if err != nil {
		return fmt.Errorf("insert newsletter: %w", err)
	}

	newsletterId := insertRes.InsertedID.(primitive.ObjectID)
	fmt.Println(newsletterId)

	recipientsCollection := ctx.mongoDB.Collection("recipients")

	filter := bson.M{"newsletterTypeId": newsletterTypeId}
	recipientsDb, err := recipientsCollection.Find(ctxBg, filter)
	if err != nil {
		return err
	}
	defer recipientsDb.Close(ctxBg)

	var recipients []RecipientsDB
	if err := recipientsDb.All(ctxBg, &recipients); err != nil {
		return fmt.Errorf("decoding recipients: %w", err)
	}

	sendGridApiKey := config.GetEnv("SENDGRID_API_KEY")
	if sendGridApiKey == "" {
		return fmt.Errorf("SENDGRID_API_KEY is not set")
	}

	sendGridClient := sendgrid.NewSendClient(sendGridApiKey)

	from := mail.NewEmail("English Daily Pill", "no.reply@takedi.com")
	subject := fmt.Sprintf("English Daily Pill — %s", info.Word)

	m := mail.NewV3Mail()
	m.SetFrom(from)
	m.Subject = subject

	m.AddContent(mail.NewContent("text/plain", getDailyWordNewsletterPlainText(info)))
	m.AddContent(mail.NewContent("text/html", getDailyWordNewsletterHTML(info)))

	for _, r := range recipients {
		p := mail.NewPersonalization()
		p.AddTos(mail.NewEmail(r.Name, r.Email))

		// If you need per-user unsubscribe/preference links, you must use a Dynamic Template
		// or SendGrid’s Subscription Tracking feature. For a quick example of subscription
		// tracking (no per-user links) you could do:
		//
		// mailSettings := mail.NewMailSettings()
		// subTrack := mail.NewSubscriptionTracking()
		// subTrack.SetEnable(true)
		// subTrack.SetText("Unsubscribe")
		// mailSettings.SetSubscriptionTracking(subTrack)
		// m.SetMailSettings(mailSettings)
		//
		// Otherwise, switch to a Dynamic Template on SendGrid, and do:
		//
		//    m.SetTemplateID("d-your-template-id")
		//    p.SetDynamicTemplateData("word",             info.Word)
		//    p.SetDynamicTemplateData("definition",       info.Definition)
		//    p.SetDynamicTemplateData("funFact",          info.FunFact)
		//    p.SetDynamicTemplateData("examples",         info.Examples)
		//    p.SetDynamicTemplateData("synonyms",         info.Synonyms)
		//    p.SetDynamicTemplateData("antonyms",         info.Antonyms)
		//    p.SetDynamicTemplateData("usageTip",         info.UsageTip)
		//    p.SetDynamicTemplateData("unsubscribe_link", fmt.Sprintf("https://…/unsub?uid=%s", r.ID.Hex()))
		//    p.SetDynamicTemplateData("preferences_link", fmt.Sprintf("https://…/prefs?uid=%s", r.ID.Hex()))

		m.AddPersonalizations(p)
	}

	respSendGrid, err := sendGridClient.Send(m)
	if err != nil {
		return fmt.Errorf("batch send error: %w", err)
	}
	if respSendGrid.StatusCode >= 400 {
		return fmt.Errorf("batch send failed: %s", respSendGrid.Body)
	}

	now = time.Now()
	update := bson.M{
		"$set": bson.M{
			"sentAt": now,
		},
	}

	_, err = dailyWordNewsletterCollection.UpdateByID(ctxBg, newsletterId, update)
	if err != nil {
		return fmt.Errorf("update newsletter: %w", err)
	}
	return nil
}
