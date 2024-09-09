package base

import(
	"flag"
)

type Command struct {
	Run func(cmd *Command,args []string)

	Long string

	Short string

	CustomFlags bool

	UsageLine string

	Flag flag.FlagSet

	Commands []*Command
}