package atlas

import (
	"context"
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/andersfylling/disgord"
)

type Event struct {
	Name string                     // the event name
	Run  func(a *Atlas) interface{} // the handler called on the event
}

var Events = make(map[string]*Event)

func NewEvent(name string, handler func(a *Atlas) interface{}) *Event {
	return &Event{name, handler}
}

// Default Atlas events
// the message handler event is used for command routing
func defaultMessageHandler(a *Atlas) interface{} {
	return func(s disgord.Session, msg *disgord.MessageCreate) {
		m := msg.Message
		prefix := a.GetPrefix(m)
		// edge case but we should be handling this
		if prefix == "" {
			return
		}

		isCommand := strings.HasPrefix(m.Content, prefix)

		if !isCommand {
			return
		}

		raw := strings.TrimLeft(m.Content, prefix)
		rawArgs := strings.Split(raw, " ")
		cmd := rawArgs[0]
		args := make(map[int]string)

		if len(rawArgs) > 1 {
			rawArgs = rawArgs[1:]

			for i, arg := range rawArgs {
				args[i] = arg
			}
		}

		command, ok := Commands[cmd]

		if !ok {
			return
		}

		ctx := Context{m, a, args, context.Background()}
		t := time.Now()
		command.Run(ctx)
		diff := int(math.Round(float64(time.Now().Sub(t).Nanoseconds() / 1e6)))

		a.Logger.Info(fmt.Sprintf("%s used command %s (%d)", m.Author.Username, command.Name, diff))
	}
}

func init() {
	Use(NewEvent("MESSAGE_CREATE", defaultMessageHandler))
}
