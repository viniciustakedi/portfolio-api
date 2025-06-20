package newslettercron

import (
	"fmt"
	"portfolio/api/api/emails"
	"time"

	"github.com/robfig/cron/v3"
)

func SendDailyPhrasalVerb() {
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

	// TODO: Get schedule time from Database
	// Schedule at 07:07:00 every day
	spec := "0 7 7 * * *"
	id, err := c.AddFunc(spec, func() {
		weekday := time.Now().In(loc).Weekday()
		if weekday != time.Tuesday && weekday != time.Thursday && weekday != time.Saturday {
			return
		}

		now := time.Now().In(loc).Format("2006-01-02T15:04:05")
		if err := emailsController.SendDailyPhrasalVerbNewsletter(); err != nil {
			fmt.Println(now, "- Error sending daily English word - Phasal Verb Edition:", err)
			return
		}
		fmt.Println(now, "- OK, Daily word sent !")
	})
	if err != nil {
		panic(fmt.Sprintf("scheduling daily word newsletter - Phrasal verb edition: %v", err))
	}

	c.Entry(id).Job.Run()

	go c.Start()
}
