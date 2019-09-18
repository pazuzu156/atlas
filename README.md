# Atlas

**This is a fork of [polaron/aurora](https://github.com/polaron/aurora)**

~~Aurora is a framework for building feature rich Discord bots in Go. Currently, Aurora only supports [Disgord](https://github.com/andersfylling/disgord), but a discordgo branch is planned.~~

Atlas is a simple command framework built for use with [Disgord](https://github.com/andersfylling/disgord)

Atlas's current goal is to take what Aurora did, and build upon it, improving it as much as I can.

Planned features are better command implementations, easier alias registrations, and help text to go along with a built-in help command.

## Installation

Atlas supports Go modules/vgo and they're the recommended way to install it. I will not be supporting installation using GOPATH.
```
module username/project

require github.com/pazuzu156/atlas v{VERSION}
```
For a list of releases, check the [tags](https://github.com/pazuzu156/atlas/releases).

## Usage

```go
// initialize atlas
client := atlas.New(&atlas.Options{
	DisgordOptions: &disgord.Config{
		BotToken: os.Getenv("TOKEN")
	}
})
```

For adding commands, check the [examples](https://github.com/pazuzu156/aurora/tree/master/examples).

## Docs

To be added
