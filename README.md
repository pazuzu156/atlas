# Aurora

Aurora is a framework for building feature rich Discord bots in Go. Currently, Aurora only supports [Disgord](https://github.com/andersfylling/disgord), but a discordgo branch is planned.

## Installation

Aurora supports Go modules/vgo and they're the recommended way to install it. I will not be supporting installation using GOPATH.
```
module username/project

require github.com/polaron/aurora v{VERSION}
```
For a list of releases, check the [tags](https://github.com/polaron/aurora/releases).

## Usage

```go
// initialize aurora
client := aurora.New(&aurora.Options{
	DisgordOptions: &disgord.Config{
		BotToken: os.Getenv("TOKEN")
	}
})
```

For adding commands, check the [examples](https://github.com/polaron/aurora/tree/master/examples).

## Docs

To be added