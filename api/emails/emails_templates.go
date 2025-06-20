package emails

import (
	"strings"
	"time"
)

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

func getDailyWordNewsletterHTML(data WordInfo) string {
	currentDate := time.Now().Format("January 2, 2006")

	html := `<!DOCTYPE html><html lang="en" xmlns="http://www.w3.org/1999/xhtml" xmlns:v="urn:schemas-microsoft-com:vml" xmlns:o="urn:schemas-microsoft-com:office:office"><head><meta charset="utf-8"><meta name="viewport" content="width=device-width,initial-scale=1.0"><meta http-equiv="X-UA-Compatible" content="IE=edge"><meta name="x-apple-disable-message-reformatting"><title>Learning with Cacau - Newsletter</title><!--[if mso]><noscript><xml><o:OfficeDocumentSettings><o:AllowPNG/><o:PixelsPerInch>96</o:PixelsPerInch></o:OfficeDocumentSettings></xml></noscript><![endif]--></head><style>.chat-bubble{position:relative;display:inline-block;margin:0;padding:15px;width:250px;background:#fff;color:#323232;font-size:24px;font-weight:500;font-family:Arial,Helvetica,sans-serif;line-height:1.2;border-radius:50px;position:absolute;top:50%;left:90%;transform:translate(-90%,-50%)}.chat-bubble::after{content:"";position:absolute;top:10px;left:0;border-width:15px 10px 0 10px;border-style:solid;rotate:10deg;border-color:#fff transparent transparent transparent}</style><body style="margin:0;padding:0;width:100%;word-break:break-word;-webkit-font-smoothing:antialiased;background-color:#f4f4f4;font-family:Arial,Helvetica,sans-serif"><div role="article" aria-roledescription="email" lang="en" style="text-size-adjust:100%;-webkit-text-size-adjust:100%;-ms-text-size-adjust:100%"><div style="display:none;max-height:0;overflow:hidden;font-size:1px;line-height:1px;color:#f4f4f4">Expand your vocabulary with today's featured word and its fascinating origins!</div><table role="presentation" cellspacing="0" cellpadding="0" border="0" width="100%" style="margin:auto;background-color:#f4f4f4"><tr><td style="padding:20px 0"><table role="presentation" cellspacing="0" cellpadding="0" border="0" width="600" style="margin:0 auto;background-color:#ffffff;border-radius:8px;box-shadow:0 2px 10px rgba(0,0,0,0.1)"><tr><td style="text-align:center;background-color:#fff;border-radius:8px 8px 0 0"><table role="presentation" cellspacing="0" cellpadding="0" border="0" width="100%"><tr><td style="text-align:center;position:relative"><h1 class="chat-bubble">Today let's learn about the word <strong>{{word}}!</strong></h1><img src="https://takedi-portfolio.s3.sa-east-1.amazonaws.com/newsletter/newsletter-daily-word.png" alt="cacau" style="width:100%;height:450px;object-fit:cover;border-radius:8px 8px 0 0"><h1 style="margin:30px 0 0 0;color:#323232;font-size:24px;font-weight:bold;font-family:Arial,Helvetica,sans-serif;line-height:1.2">Learn with Cacau, woof woof!</h1><p style="margin:5px 0 30px 0;color:#323232;font-size:16px;font-family:Arial,Helvetica,sans-serif">{{current_date}}</p></td></tr></table></td></tr><tr><td style="padding:20px 20px 15px 20px;text-align:center;background-color:#68a5d438"><h2 style="margin:0 0 10px 0;color:#0D0D0D;font-size:24px;font-weight:bold;font-family:Arial,Helvetica,sans-serif;line-height:1.1">Our daily word is {{word}}</h2></td></tr><tr><td style="padding:0 40px 20px 40px"><h3 style="margin:30px 0 15px 0;color:#0D0D0D;font-size:18px;font-weight:bold;font-family:Arial,Helvetica,sans-serif">Definition</h3><p style="margin:0 0 20px 0;color:#555555;font-size:16px;line-height:1.6;font-family:Arial,Helvetica,sans-serif">{{definition}}</p><div style="background-color:#68A5D4;border-left:4px solid #235888;padding:15px 20px;margin:25px 0;border-radius:0 4px 4px 0"><p style="margin:0;color:#fff;font-size:16px;font-family:Arial,Helvetica,sans-serif"><strong>Did you know?</strong> {{funFact}}</p></div></td></tr><tr><td style="padding:0 40px 30px 40px"><h3 style="margin:0 0 15px 0;color:#333333;font-size:18px;font-weight:bold;font-family:Arial,Helvetica,sans-serif">Example Sentences</h3>{{examples}}</td></tr><tr><td style="padding:0 40px 20px 40px;margin:50px 0 0 0"><table role="presentation" cellspacing="0" cellpadding="0" border="0" width="100%"><tr><td style="width:50%;vertical-align:top;padding-right:15px;margin:50px 0 0 0"><h4 style="margin:0 0 10px 0;color:#333333;font-size:18px;font-weight:bold;font-family:Arial,Helvetica,sans-serif">Synonyms</h4><p style="margin:0;color:#555555;font-size:16px;line-height:1.5;font-family:Arial,Helvetica,sans-serif">{{synonyms}}</p></td><td style="width:50%;vertical-align:top;padding-left:15px"><h4 style="margin:0 0 10px 0;color:#333333;font-size:18px;font-weight:bold;font-family:Arial,Helvetica,sans-serif">Antonyms</h4><p style="margin:0;color:#555555;font-size:16px;line-height:1.5;font-family:Arial,Helvetica,sans-serif">{{antonyms}}</p></td></tr></table><div style="background-color:#68A5D4;border-left:4px solid #235888;padding:15px 20px;margin:30px 0;border-radius:0 4px 4px 0"><p style="margin:0;color:#fff;font-size:16px;font-family:Arial,Helvetica,sans-serif"><strong>Usage Tip:</strong> {{usageTip}}</p></div></td></tr><tr><td style="padding:30px 40px;background-color:#f8f9fa;border-top:1px solid #e9ecef"><table role="presentation" cellspacing="0" cellpadding="0" border="0" width="100%"><tr><td style="text-align:center"><p style="margin:0 0 15px 0;color:#888888;font-size:12px;font-family:Arial,Helvetica,sans-serif">By Vinicius Takedi<br> São Paulo, Brazil.</p><p style="margin:0;color:#888888;font-size:11px;font-family:Arial,Helvetica,sans-serif">You're receiving this email because you subscribed to our Word of the Day newsletter.<br><a href="{{unsubscribe_link}}" style="color:#888888;text-decoration:underline">Unsubscribe</a> | <a href="{{preferences_link}}" style="color:#888888;text-decoration:underline">Update preferences</a></p></td></tr></table></td></tr></table></td></tr></table></div></body></html>`
	html = strings.Replace(html, "{{current_date}}", currentDate, -1)
	html = strings.Replace(html, "{{word}}", data.Word, -1)
	html = strings.Replace(html, "{{definition}}", data.Definition, -1)
	html = strings.Replace(html, "{{funFact}}", data.FunFact, -1)
	html = strings.Replace(html, "{{usageTip}}", data.UsageTip, -1)

	var exBuilder strings.Builder
	for _, ex := range data.Examples {
		exBuilder.WriteString(`<div style="margin-bottom: 15px;">`)
		exBuilder.WriteString("\n<p style=\"margin: 0; color: #555555; font-size: 18px; line-height: 1.6; font-family: Arial, Helvetica, sans-serif;\">")
		exBuilder.WriteString(ex)
		exBuilder.WriteString("</p>\n</div>\n")
	}
	html = strings.Replace(html, "{{examples}}", exBuilder.String(), -1)

	html = strings.Replace(html, "{{synonyms}}", strings.Join(data.Synonyms, ", "), -1)
	html = strings.Replace(html, "{{antonyms}}", strings.Join(data.Antonyms, ", "), -1)

	return html
}

func getDailyWordNewsletterPlainText(data WordInfo) string {
	loc, err := time.LoadLocation("America/Sao_Paulo")
	if err != nil {
		loc = time.Local
	}
	currentDate := time.Now().In(loc).Format("January 2, 2006")

	var sb strings.Builder

	sb.WriteString("Become smarter with Cacau\n")
	sb.WriteString(currentDate + "\n\n")

	sb.WriteString("Phrasal verb of the Day: " + data.Word + "\n\n")

	sb.WriteString("Definition:\n")
	sb.WriteString(data.Definition + "\n\n")

	sb.WriteString("Did you know?\n")
	sb.WriteString(data.FunFact + "\n\n")

	sb.WriteString("Examples:\n")
	for _, ex := range data.Examples {
		sb.WriteString("- " + ex + "\n")
	}
	sb.WriteString("\n")

	sb.WriteString("Synonyms: " + strings.Join(data.Synonyms, ", ") + "\n")
	sb.WriteString("Antonyms: " + strings.Join(data.Antonyms, ", ") + "\n\n")

	sb.WriteString("Usage Tip:\n")
	sb.WriteString(data.UsageTip + "\n\n")

	sb.WriteString("By Vinicius Takedi\n")
	sb.WriteString("São Paulo, Brazil.\n\n")

	sb.WriteString("You're receiving this email because you subscribed to our Lern with Cacau newsletter.\n")
	sb.WriteString("Unsubscribe: {{unsubscribe_link}}\n")
	sb.WriteString("Update preferences: {{preferences_link}}\n")

	return sb.String()
}

func getDailyPhrasalVerbNewsletterHTML(data WordInfo) string {
	currentDate := time.Now().Format("January 2, 2006")

	html := `<!DOCTYPE html><html lang="en" xmlns="http://www.w3.org/1999/xhtml" xmlns:v="urn:schemas-microsoft-com:vml" xmlns:o="urn:schemas-microsoft-com:office:office"><head><meta charset="utf-8"><meta name="viewport" content="width=device-width,initial-scale=1.0"><meta http-equiv="X-UA-Compatible" content="IE=edge"><meta name="x-apple-disable-message-reformatting"><title>Learning with Cacau - Newsletter</title><!--[if mso]><noscript><xml><o:OfficeDocumentSettings><o:AllowPNG/><o:PixelsPerInch>96</o:PixelsPerInch></o:OfficeDocumentSettings></xml></noscript><![endif]--></head><style>.chat-bubble{position:relative;display:inline-block;margin:0;padding:15px;width:250px;background:#fff;color:#323232;font-size:24px;font-weight:500;font-family:Arial,Helvetica,sans-serif;line-height:1.2;border-radius:50px;position:absolute;top:50%;left:90%;transform:translate(-90%,-50%)}.chat-bubble::after{content:"";position:absolute;top:10px;left:0;border-width:15px 10px 0 10px;border-style:solid;rotate:10deg;border-color:#fff transparent transparent transparent}</style><body style="margin:0;padding:0;width:100%;word-break:break-word;-webkit-font-smoothing:antialiased;background-color:#f4f4f4;font-family:Arial,Helvetica,sans-serif"><div role="article" aria-roledescription="email" lang="en" style="text-size-adjust:100%;-webkit-text-size-adjust:100%;-ms-text-size-adjust:100%"><div style="display:none;max-height:0;overflow:hidden;font-size:1px;line-height:1px;color:#f4f4f4">Expand your vocabulary with today's featured phrasal verb and its fascinating origins!</div><table role="presentation" cellspacing="0" cellpadding="0" border="0" width="100%" style="margin:auto;background-color:#f4f4f4"><tr><td style="padding:20px 0"><table role="presentation" cellspacing="0" cellpadding="0" border="0" width="600" style="margin:0 auto;background-color:#ffffff;border-radius:8px;box-shadow:0 2px 10px rgba(0,0,0,0.1)"><tr><td style="text-align:center;background-color:#fff;border-radius:8px 8px 0 0"><table role="presentation" cellspacing="0" cellpadding="0" border="0" width="100%"><tr><td style="text-align:center;position:relative"><h1 class="chat-bubble">Today let's learn about the phrasal verb <strong>{{word}}!</strong></h1><img src="https://takedi-portfolio.s3.sa-east-1.amazonaws.com/newsletter/newsletter-daily-word.png" alt="cacau" style="width:100%;height:450px;object-fit:cover;border-radius:8px 8px 0 0"><h1 style="margin:30px 0 0 0;color:#323232;font-size:24px;font-weight:bold;font-family:Arial,Helvetica,sans-serif;line-height:1.2">Learn with Cacau, woof woof!</h1><p style="margin:5px 0 30px 0;color:#323232;font-size:16px;font-family:Arial,Helvetica,sans-serif">{{current_date}}</p></td></tr></table></td></tr><tr><td style="padding:20px 20px 15px 20px;text-align:center;background-color:#68a5d438"><h2 style="margin:0 0 10px 0;color:#0D0D0D;font-size:24px;font-weight:bold;font-family:Arial,Helvetica,sans-serif;line-height:1.1">Our phrasal verb for today is {{word}}</h2></td></tr><tr><td style="padding:0 40px 20px 40px"><h3 style="margin:30px 0 15px 0;color:#0D0D0D;font-size:18px;font-weight:bold;font-family:Arial,Helvetica,sans-serif">Definition</h3><p style="margin:0 0 20px 0;color:#555555;font-size:16px;line-height:1.6;font-family:Arial,Helvetica,sans-serif">{{definition}}</p><div style="background-color:#68A5D4;border-left:4px solid #235888;padding:15px 20px;margin:25px 0;border-radius:0 4px 4px 0"><p style="margin:0;color:#fff;font-size:16px;font-family:Arial,Helvetica,sans-serif"><strong>Did you know?</strong> {{funFact}}</p></div></td></tr><tr><td style="padding:0 40px 30px 40px"><h3 style="margin:0 0 15px 0;color:#333333;font-size:18px;font-weight:bold;font-family:Arial,Helvetica,sans-serif">Example Sentences</h3>{{examples}}</td></tr><tr><td style="padding:0 40px 20px 40px;margin:50px 0 0 0"><table role="presentation" cellspacing="0" cellpadding="0" border="0" width="100%"><tr><td style="width:50%;vertical-align:top;padding-right:15px;margin:50px 0 0 0"><h4 style="margin:0 0 10px 0;color:#333333;font-size:18px;font-weight:bold;font-family:Arial,Helvetica,sans-serif">Synonyms</h4><p style="margin:0;color:#555555;font-size:16px;line-height:1.5;font-family:Arial,Helvetica,sans-serif">{{synonyms}}</p></td><td style="width:50%;vertical-align:top;padding-left:15px"><h4 style="margin:0 0 10px 0;color:#333333;font-size:18px;font-weight:bold;font-family:Arial,Helvetica,sans-serif">Antonyms</h4><p style="margin:0;color:#555555;font-size:16px;line-height:1.5;font-family:Arial,Helvetica,sans-serif">{{antonyms}}</p></td></tr></table><div style="background-color:#68A5D4;border-left:4px solid #235888;padding:15px 20px;margin:30px 0;border-radius:0 4px 4px 0"><p style="margin:0;color:#fff;font-size:16px;font-family:Arial,Helvetica,sans-serif"><strong>Usage Tip:</strong> {{usageTip}}</p></div></td></tr><tr><td style="padding:30px 40px;background-color:#f8f9fa;border-top:1px solid #e9ecef"><table role="presentation" cellspacing="0" cellpadding="0" border="0" width="100%"><tr><td style="text-align:center"><p style="margin:0 0 15px 0;color:#888888;font-size:12px;font-family:Arial,Helvetica,sans-serif">By Vinicius Takedi<br> São Paulo, Brazil.</p><p style="margin:0;color:#888888;font-size:11px;font-family:Arial,Helvetica,sans-serif">You're receiving this email because you subscribed to our Learn with Cacau newsletter.<br><a href="{{unsubscribe_link}}" style="color:#888888;text-decoration:underline">Unsubscribe</a> | <a href="{{preferences_link}}" style="color:#888888;text-decoration:underline">Update preferences</a></p></td></tr></table></td></tr></table></td></tr></table></div></body></html>`
	html = strings.Replace(html, "{{current_date}}", currentDate, -1)
	html = strings.Replace(html, "{{word}}", data.Word, -1)
	html = strings.Replace(html, "{{definition}}", data.Definition, -1)
	html = strings.Replace(html, "{{funFact}}", data.FunFact, -1)
	html = strings.Replace(html, "{{usageTip}}", data.UsageTip, -1)

	var exBuilder strings.Builder
	for _, ex := range data.Examples {
		exBuilder.WriteString(`<div style="margin-bottom: 15px;">`)
		exBuilder.WriteString("\n<p style=\"margin: 0; color: #555555; font-size: 18px; line-height: 1.6; font-family: Arial, Helvetica, sans-serif;\">")
		exBuilder.WriteString(ex)
		exBuilder.WriteString("</p>\n</div>\n")
	}
	html = strings.Replace(html, "{{examples}}", exBuilder.String(), -1)

	html = strings.Replace(html, "{{synonyms}}", strings.Join(data.Synonyms, ", "), -1)
	html = strings.Replace(html, "{{antonyms}}", strings.Join(data.Antonyms, ", "), -1)

	return html
}

func getDailyPhrasalVerbNewsletterPlainText(data WordInfo) string {
	loc, err := time.LoadLocation("America/Sao_Paulo")
	if err != nil {
		loc = time.Local
	}
	currentDate := time.Now().In(loc).Format("January 2, 2006")

	var sb strings.Builder

	sb.WriteString("Become smarter with Cacau\n")
	sb.WriteString(currentDate + "\n\n")

	sb.WriteString("Phrasal verb of the Day: " + data.Word + "\n\n")

	sb.WriteString("Definition:\n")
	sb.WriteString(data.Definition + "\n\n")

	sb.WriteString("Did you know?\n")
	sb.WriteString(data.FunFact + "\n\n")

	sb.WriteString("Examples:\n")
	for _, ex := range data.Examples {
		sb.WriteString("- " + ex + "\n")
	}
	sb.WriteString("\n")

	sb.WriteString("Synonyms: " + strings.Join(data.Synonyms, ", ") + "\n")
	sb.WriteString("Antonyms: " + strings.Join(data.Antonyms, ", ") + "\n\n")

	sb.WriteString("Usage Tip:\n")
	sb.WriteString(data.UsageTip + "\n\n")

	sb.WriteString("By Vinicius Takedi\n")
	sb.WriteString("São Paulo, Brazil.\n\n")

	sb.WriteString("You're receiving this email because you subscribed to our Lern with Cacau newsletter.\n")
	sb.WriteString("Unsubscribe: {{unsubscribe_link}}\n")
	sb.WriteString("Update preferences: {{preferences_link}}\n")

	return sb.String()
}
