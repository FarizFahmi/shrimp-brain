package logger

import (
	"fmt"
	"strings"
)

func (l *Log) Fatal(message ...any) {
	var msg strings.Builder
	if len(message) > 0 {
		msg.WriteString("[FATAL] ")
		msg.WriteString(fmt.Sprintf("%v", message[0]))
		message[0] = msg.String()
	}
	l.newLog.Fatal().Str("category", l.flag).Msg(fmt.Sprint(message...))
}
