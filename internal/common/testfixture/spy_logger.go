package testfixture

import "fmt"

type SpyLogger struct {
	Logs []string
}

func NewSpyLogger() *SpyLogger {
	return &SpyLogger{
		Logs: make([]string, 0),
	}
}

func (lg *SpyLogger) Println(v ...interface{}) {
	lg.Logs = append(lg.Logs, fmt.Sprintln(v...))
}
