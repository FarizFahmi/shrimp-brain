package logger

import (
	"os"
	"strings"

	"github.com/rs/zerolog"
)

type Log struct {
	// stdout *log.Logger
	// stderr *log.Logger
	newLog *zerolog.Logger
	flag   string
}

// func New(prefix string) *Log {
// 	// now := time.Now().Format("2006/01/02 15:04:05")
// 	return &Log{
// 		// stdout: log.New(os.Stdout, "["+now+"] [LOG]["+prefix+"]", log.Lshortfile),
// 		// stderr: log.New(os.Stderr, "["+now+"] [ERROR]["+prefix+"]", log.Lshortfile),
// 		stdout: log.New(os.Stdout, "[LOG]["+prefix+"]", log.Ldate|log.Ltime),
// 		stderr: log.New(os.Stderr, "[ERROR]["+prefix+"]", log.Ldate|log.Ltime),
// 	}
// }

// NewLogger creates and configures a new logger
func New(prefix string) *Log {
	// switch mode {
	// case zerolog.LevelDebugValue:
	// 	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	// default:
	// 	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	// }

	// Set a human-readable time format (RFC3339)
	zerolog.TimeFieldFormat = "2006/01/02 15:04:05"

	// Customize the log output
	output := zerolog.ConsoleWriter{
		Out:           os.Stdout,
		TimeFormat:    "2006/01/02 15:04:05",
		PartsOrder:    []string{"time", "category", "message"},
		FieldsExclude: []string{"category"},
		FormatLevel: func(i interface{}) string {
			return "[" + strings.ToUpper(i.(string)) + "]"
		},
		FormatTimestamp: func(i interface{}) string {
			return "[" + i.(string) + "]"
		},
		FormatFieldName: func(i interface{}) string {
			return ""
		},
		FormatFieldValue: func(i interface{}) string {
			// Only include the category if it's not empty
			if category, ok := i.(string); ok && category != "" {
				return "[" + category + "]" // Include brackets for the category
			}
			return ""
		},
		FormatMessage: func(i interface{}) string {
			return i.(string)
		},
	}

	// Create a base logger and return it as a pointer
	logger := zerolog.New(output).With().Timestamp().Logger()
	// return &logger // Return a pointer to the logger

	return &Log{
		newLog: &logger,
		flag:   prefix,
	}
}