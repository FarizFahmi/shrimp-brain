package logger

import (
	"fmt"
	"strings"
)

func (l *Log) Panic(message ...any) {
	var msg strings.Builder
	if len(message) > 0 {
		msg.WriteString("[PANIC] ")
		msg.WriteString(fmt.Sprintf("%v", message[0]))
		message[0] = msg.String()
	}
	l.newLog.Panic().Str("category", l.flag).Msg(fmt.Sprint(message...))
}
