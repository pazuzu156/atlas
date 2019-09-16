# Aurora

**This is a fork of [polaron/aurora](https://github.com/polaron/aurora)**

~~Aurora is a framework for building feature rich Discord bots in Go. Currently, Aurora only supports [Disgord](https://github.com/andersfylling/disgord), but a discordgo branch is planned.~~

Aurora is a simple command framework built for use with [Disgord](https://github.com/andersfylling/disgord)

## Installation

Aurora supports Go modules/vgo and they're the recommended way to install it. I will not be supporting installation using GOPATH.
```
module username/project

require github.com/pazuzu156/aurora v{VERSION}
```
For a list of releases, check the [tags](https://github.com/pazuzu156/aurora/releases).

## Usage

```go
// initialize aurora
client := aurora.New(&aurora.Options{
	DisgordOptions: &disgord.Config{
		BotToken: os.Getenv("TOKEN")
	}
})
```

For adding commands, check the [examples](https://github.com/pazuzu156/aurora/tree/master/examples).

## Docs

To be added
