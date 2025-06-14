package newslettercron

import (
	"fmt"
	"portfolio/api/api/emails"
	"time"

	"github.com/robfig/cron/v3"
)

func SendDailyWord() {
	// Determine GMT-3 (São Paulo) location
	loc, err := time.LoadLocation("America/Sao_Paulo")
	if err != nil {
		loc = time.FixedZone("GMT-3", -3*3600)
	}

	fmt.Println(time.Now().In(loc).Format("2006-01-02T15:04:05"), "- Cron to send the daily English word")
	emailsController := emails.MakeEmailsController()

	// Create a cron with seconds and the right TZ
	c := cron.New(
		cron.WithSeconds(),
		cron.WithLocation(loc),
	)

	// Schedule at 07:07:00 every day
	spec := "0 7 7 * * *"
	_, err = c.AddFunc(spec, func() {
		now := time.Now().In(loc).Format("2006-01-02T15:04:05")
		if err := emailsController.SendDailyWordNewsletter(); err != nil {
			fmt.Println(now, "- Error sending daily English word:", err)
			return
		}
		fmt.Println(now, "- OK, sent!")
	})
	if err != nil {
		panic(fmt.Sprintf("scheduling daily word newsletter: %v", err))
	}

	// c.Entry(id).Job.Run()

	go c.Start()
}
