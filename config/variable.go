package config

import (
	"announce/utils"
	"strconv"
)

var (
	Mode                 = utils.MustGetEnv("MODE")
	SpaceNotif           = utils.MustGetEnv("SPACE_NOTIF")
	GrcNotif             = utils.MustGetEnv("GRC_NOTIF")
	TimeScheduleNotif, _ = strconv.Atoi(utils.MustGetEnv("TIME_SCHEDULE_NOTIF"))

	// Ninja API
	NinjaApiBaseUrl   = utils.MustGetEnv("NINJA_API_BASE_URL")
	NinjaApiKey       = utils.MustGetEnv("NINJA_API_KEY")
	NinjaApiQuotesUrl = utils.MustGetEnv("NINJA_API_QUOTES_URL")
	NinjaApiJokesUrl  = utils.MustGetEnv("NINJA_API_JOKES_URL")
	NinjaApiRiddleUrl = utils.MustGetEnv("NINJA_API_RIDDLE_URL")
	NinjaApiTriviaUrl = utils.MustGetEnv("NINJA_API_TRIVIA_URL")
	NinjaApiAdviceUrl = utils.MustGetEnv("NINJA_API_ADVICE_URL")
	NinjaApiFactUrl   = utils.MustGetEnv("NINJA_API_FACT_URL")

	KyAttendanceUrl = utils.MustGetEnv("KY_ATTENDANCE_URL")
)
