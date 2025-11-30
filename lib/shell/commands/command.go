package commands

type Command interface {
	GetHelp() string
	Run(args []string) error
}
