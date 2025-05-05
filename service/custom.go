package service

import (
	"announce/constant"
	"announce/helper"
)

func GRC() {
	// For testing purposes
	// now := time.Date(nowA.Year(), nowA.Month()+2, 5, 0, 0, 0, 0, nowA.Local().Location())

	if err := helper.HandleSendToSpace(constant.Notif["GRC"], "🔔 [REMINDER] 🔔 \n", "JANGAN LUPA UPDATE PROGRESS !!"); err != nil {
		log.Log(err.Error())
	}
}
