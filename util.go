package aurora

import "log"

type logger struct{}

func (l *logger) Info(v ...interface{}) {
	log.Print("info: ", v)
}

func (l *logger) Debug(v ...interface{}) {
	log.Print("debug: ", v)
}

func (l *logger) Error(v ...interface{}) {
	log.Print("error: ", v)
}

func DefaultLogger() Logger {
	return &logger{}
}

func init() {
	log.SetPrefix("[AURORA] ")
}
