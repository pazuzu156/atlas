package main

import (
	"github.com/andersfylling/disgord"
	"github.com/pazuzu156/atlas"
	
	"os"
)

var pingCommand = atlas.NewCommand("ping").SetDescription("Ping/pong command")

func main() {
	client := atlas.New(&atlas.Options{
		DisgordOptions: &disgord.Config{
			BotToken: os.Getenv("DISCORD_TOKEN"),
			Logger:   disgord.DefaultLogger(false),
		},
	})

	client.Use(atlas.DefaultLogger())
	client.GetPrefix = func(m *disgord.Message) string {
		return "!"
	}
	if err := client.Init(); err != nil {
		panic(err)
	}
}

func init() {
	pingCommand.Run = func(c atlas.Context) {
		c.Message.RespondString(c.Atlas, "Pong!")
	}

	atlas.Use(pingCommand)
}
