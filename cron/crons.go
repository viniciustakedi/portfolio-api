package cron

import newslettercron "portfolio/api/cron/newsletters"

func Init() {
	newslettercron.SendDailyWord()
	newslettercron.SendDailyPhrasalVerb()
}
