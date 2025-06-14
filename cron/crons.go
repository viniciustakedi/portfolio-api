package cron

import newslettercron "portfolio/api/cron/newsletters"

func Init() {
	newslettercron.SendDailyWord()
	// Add more cron jobs here as needed
}
