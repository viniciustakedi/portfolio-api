package emails

import "strings"

func getPortfolioMessageHTML(data SendPortfolioMessage) string {
	html := `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8" /><title>New portfolio message!</title><style>body {font-family: Arial, sans-serif;background-color: #f9f9f9;padding: 20px;margin: 0;}.email-container {max-width: 500px;margin: 0 auto;background-color: #ffffff;border-radius: 10px;padding: 40px 20px;text-align: center;box-shadow: 0 2px 8px rgba(0, 0, 0, 0.05);}.emoji {font-size: 48px;margin-bottom: 20px;}h1 {font-size: 22px;margin-bottom: 10px;}p {font-size: 16px;color: #555555;margin: 10px 0;}a.button {display: inline-block;margin-top: 20px;padding: 12px 24px;background-color: #007bff;color: #ffffff;text-decoration: none;border-radius: 6px;font-weight: bold;}.footer {font-size: 12px;color: #999999;margin-top: 30px;}</style></head><body><div class="email-container"><div class="emoji">🥳</div><h1>{{sender_name}} sent you a message!</h1><p>To reply {{sender_name}} use your professional email <strong>contact@takedi.com</strong> and send a message to <strong>{{sender_email}}</strong>.</p><p style="margin-top: 30px;background-color: #f1f1f1;border-radius: 8px;padding: 15px;">"{{sender_message}}"</p><a href="https://takedi.com" class="button">Click here to go to portfolio</a><p class="footer">This email was sent automatically to Vinicius Takedi. If you received this by mistake, please disregard it. Thank you very much!</p></div></body></html>`

	html = strings.Replace(html, "{{sender_name}}", data.Name, -1)
	html = strings.Replace(html, "{{sender_email}}", data.Email, -1)
	html = strings.Replace(html, "{{sender_message}}", data.Message, -1)

	return html
}

func getPortfolioMessagePlainText(data SendPortfolioMessage) string {
	plaintext := `🥳 New portfolio message! {{sender_name}} sent you a message! To reply {{sender_name}}, use your professional email contact@takedi.com and send a message to {{sender_email}}. Message: "{{sender_message}}" Click here to go to portfolio website: https://takedi.com --- This email was sent automatically to Vinicius Takedi. If you received this by mistake, please disregard it. Thank you very much!`

	plaintext = strings.Replace(plaintext, "{{sender_name}}", data.Name, -1)
	plaintext = strings.Replace(plaintext, "{{sender_email}}", data.Email, -1)
	plaintext = strings.Replace(plaintext, "{{sender_message}}", data.Message, -1)

	return plaintext
}
