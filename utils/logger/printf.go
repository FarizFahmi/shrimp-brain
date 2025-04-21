package logger

import "fmt"

func (l *Log) Printf(apaini string, message ...any) {
	fmt.Println("HEHEHEHE : " + apaini)
	if len(message) > 0 {
		message[0] = "📢 " + fmt.Sprintf("%v", message[0])
	}
	// l.stdout.Println(message...)
	l.newLog.Info().Str("category", l.flag).Msg(fmt.Sprint(message...))
}
