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

	fmt.Println(time.Now().In(loc).Format("2006-01-02T15:04:05"), "- Cron to send the daily English word - Word Edition")
	emailsController := emails.MakeEmailsController()

	// Create a cron with seconds and the right TZ
	c := cron.New(
		cron.WithSeconds(),
		cron.WithLocation(loc),
	)

	allowedWeekdays := map[time.Weekday]bool{
		time.Monday:    true,
		time.Wednesday: true,
		time.Friday:    true,
		time.Sunday:    true,
	}

	spec, err := emailsController.GetNewsletterScheduleTime("684cd13895298f80e21813a9")
	if err != nil {
		panic(fmt.Sprintf("Erro to get schedule for learn with cacau - word mode: %v", err))
	}

	fmt.Println(
		time.Now().In(loc).Format("2006-01-02T15:04:05"),
		"- Cron to send the daily English word - Word Edition",
		fmt.Sprintf(" - Scheduled to %s", spec),
	)

	_, err = c.AddFunc(spec, func() {
		weekday := time.Now().In(loc).Weekday()

		if !allowedWeekdays[weekday] {
			return
		}

		now := time.Now().In(loc).Format("2006-01-02T15:04:05")
		if err := emailsController.SendDailyWordNewsletter(); err != nil {
			fmt.Println(now, "- Error sending daily English word:", err)
			return
		}
		fmt.Println(now, "- OK, Daily word sent !")
	})
	if err != nil {
		panic(fmt.Sprintf("scheduling daily word newsletter: %v", err))
	}

	// c.Entry(id).Job.Run()
	go c.Start()
}
