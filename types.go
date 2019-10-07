package atlas

import "github.com/andersfylling/disgord"

type Disgord interface {
	Connect() error
	DisconnectOnInterrupt() error
	disgord.Session
}

type Atlas struct {
	Disgord
	Options   *Options
	Logger    Logger
	GetPrefix func(m *disgord.Message) string
}

type Options struct {
	OwnerID        string
	DisgordOptions *disgord.Config
}

type Context struct {
	Message *disgord.Message
	Atlas   *Atlas
	Args    map[int]string
}

type Logger interface {
	Info(v ...interface{})
	Debug(v ...interface{})
	Error(v ...interface{})
}
