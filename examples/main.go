package main

import (
	"github.com/andersfylling/disgord"
	"github.com/pazuzu156/aurora"
	
	"os"
)

var pingCommand = aurora.NewCommand("ping").SetDescription("Ping/pong command")

func main() {
	client := aurora.New(&aurora.Options{
		DisgordOptions: &disgord.Config{
			BotToken: os.Getenv("DISCORD_TOKEN"),
			Logger:   disgord.DefaultLogger(false),
		},
	})

	client.Use(aurora.DefaultLogger())
	client.GetPrefix = func(m *disgord.Message) string {
		return "!"
	}
	if err := client.Init(); err != nil {
		panic(err)
	}
}

func init() {
	pingCommand.Run = func(c aurora.Context) {
		c.Message.RespondString(c.Aurora, "Pong!")
	}

	aurora.Use(pingCommand)
}
