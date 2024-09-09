package main

import (
	"github.com/chengshanshan/v2ray-core/main/commands"
	"github.com/chengshanshan/v2ray-core/main/commands/base"
)

func main(){
	base.RootCommand.Long="A unified platform for anti-censorship."
	base.RegisterCommand(Commands.CmdRun)
	
}