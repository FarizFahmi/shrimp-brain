package logger

import (
	"fmt"
	"strings"
)

func (l *Log) Log(message ...any) {
	var msg strings.Builder
	if len(message) > 0 {
		msg.WriteString("[INFO] ")
		msg.WriteString(fmt.Sprintf("%v", message[0]))
		message[0] = msg.String()
	}
	
	l.newLog.Info().Str("category", l.flag).Msg(fmt.Sprint(message...))
}
