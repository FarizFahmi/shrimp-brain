package logger

import (
	"fmt"
	"runtime/debug"
	"strings"
)

func (l *Log) Error(message ...any) {
	var msg strings.Builder
	if len(message) > 0 {
		msg.WriteString("[ERROR] ")
		msg.WriteString(fmt.Sprintf("%v", message[0]))
		message[0] = msg.String()

		message = append(message, "\n\n", string(debug.Stack()))
	}
	l.newLog.Error().Str("category", l.flag).Msg(fmt.Sprint(message...))
}

func (l *Log) ErrorWithoutTrace(message ...any) {
	var msg strings.Builder
	if len(message) > 0 {
		msg.WriteString("[ERROR] ")
		msg.WriteString(fmt.Sprintf("%v", message[0]))
		message[0] = msg.String()
	}
	l.newLog.Info().Str("category", l.flag).Msg(fmt.Sprint(message...))
}
