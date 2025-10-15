package app

import "github.com/fkl13/pcla/distributing/notify"

func send_notification(msg string) {
	n := notify.New("Pomodoro", msg, notify.SeverityNormal)
	n.Send()
}
