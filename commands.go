package aurora

var Commands = make(map[string]*Command)

type Command struct {
	Name         string            // the name of the command
	Aliases      []string          // other valid command names
	Description  string            // a description of the command, shown in the help message
	ExtendedHelp string            // extended help
	Run          func(ctx Context) // the code called when the command is executed
}

func NewCommand(name string) *Command {
	return &Command{Name: name}
}

func (c *Command) SetAliases(aliases ...string) *Command {
	c.Aliases = append(c.Aliases, aliases...)
	return c
}

func (c *Command) SetDescription(description string) *Command {
	c.Description = description
	return c
}
