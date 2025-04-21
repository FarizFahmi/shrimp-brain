package service

import (
	"announce/constant"
	"announce/helper"
	"fmt"
	"time"
)

func getWeekNumberInMonth(date time.Time) int {
	log.Log("Date: ", date)

	// Find the first Thursday of the month
	firstDayOfMonth := time.Date(date.Year(), date.Month(), 1, 0, 0, 0, 0, date.Location())
	offset := (time.Thursday - firstDayOfMonth.Weekday() + 7) % 7
	firstThursday := firstDayOfMonth.AddDate(0, 0, int(offset))

	// Add all Thursdays of the month
	var countThursdays int
	for thursday := firstThursday; thursday.Month() == date.Month(); thursday = thursday.AddDate(0, 0, 7) {
		countThursdays++
		if date.Day() == thursday.Day() {
			break
		}
	}

	log.Log("thursday : ", countThursdays)

	return countThursdays
}

func Notif() {
	nowA := time.Now()

	// For testing purposes
	// now := time.Date(nowA.Year(), nowA.Month()+2, 5, 0, 0, 0, 0, nowA.Local().Location())
	
	log.Log("Now : ", nowA)
	thursNum := getWeekNumberInMonth(nowA)

	uniform := "HITAM"
	if thursNum%2 == 0 {
		uniform = "PUTIH"
	}

	if err := helper.HandleSendToSpace(constant.Notif["UNIFORM"], fmt.Sprintf("🔔 [SERAGAM-KAMIS KE-%d] 🔔 \n", thursNum), fmt.Sprintf("Moshi², jangan lupa hari ini pakai seragam warna <b>%s</b> ya", uniform)); err != nil {
		log.Log(err.Error())
	}
}
