package main

import (
	"announce/config"
	"announce/service"
	"announce/utils/logger"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-co-op/gocron/v2"
)

var log = logger.New("MAIN")

func main() {
	sch, err := gocron.NewScheduler()
	if err != nil {
		log.Log(err.Error())
	}

	if config.Mode == "test" {
		log.Log("Scheduler test started")

		// uniform
		if _, err = sch.NewJob(
			gocron.DurationJob(time.Duration(config.TimeScheduleNotif)*time.Second),
			gocron.NewTask(service.Notif),
		); err != nil {
			log.Log(err.Error())
		}

		// checkin
		if _, err = sch.NewJob(
			gocron.DurationJob(time.Duration(config.TimeScheduleNotif)*time.Second),
			gocron.NewTask(service.HandleAttendance),
		); err != nil {
			log.Log(err.Error())
		}

		// checkout
		if _, err = sch.NewJob(
			gocron.DurationJob(time.Duration(config.TimeScheduleNotif)*time.Second),
			gocron.NewTask(service.HandleCheckout),
		); err != nil {
			log.Log(err.Error())
		}

		// GRC
		if _, err = sch.NewJob(
			gocron.DurationJob(time.Duration(config.TimeScheduleNotif)*time.Second),
			gocron.NewTask(service.GRC),
		); err != nil {
			log.Log(err.Error())
		}

	} else {
		log.Log("Scheduler started")

		// uniform
		if _, err = sch.NewJob(
			gocron.WeeklyJob(1, gocron.NewWeekdays(time.Thursday), gocron.NewAtTimes(gocron.NewAtTime(5, 0, 0))),
			gocron.NewTask(service.Notif),
		); err != nil {
			log.Log(err.Error())
		}

		// checkin
		if _, err = sch.NewJob(
			gocron.WeeklyJob(1, gocron.NewWeekdays(time.Monday, time.Tuesday, time.Wednesday, time.Thursday, time.Friday), gocron.NewAtTimes(gocron.NewAtTime(8, 57, 0))),
			gocron.NewTask(service.HandleAttendance),
		); err != nil {
			log.Log(err.Error())
		}

		// checkout
		if _, err = sch.NewJob(
			gocron.WeeklyJob(1, gocron.NewWeekdays(time.Monday, time.Tuesday, time.Wednesday, time.Thursday, time.Friday), gocron.NewAtTimes(gocron.NewAtTime(17, 57, 0))),
			gocron.NewTask(service.HandleCheckout),
		); err != nil {
			log.Log(err.Error())
		}

		// GRC
		if _, err = sch.NewJob(
			gocron.WeeklyJob(1, gocron.NewWeekdays(time.Monday, time.Tuesday, time.Wednesday, time.Thursday, time.Friday), gocron.NewAtTimes(gocron.NewAtTime(9, 0, 0))),
			gocron.NewTask(service.GRC),
		); err != nil {
			log.Log(err.Error())
		}
	}

	sch.Start()

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, os.Interrupt, syscall.SIGTERM)

	<-sigchan
	log.Log("Shutting down scheduler\n")
	if err := sch.Shutdown(); err != nil {
		log.Log(err.Error())
	}
}
