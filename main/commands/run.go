package Commands

import (
	core "github.com/chengshanshan/v2ray-core"
	"github.com/chengshanshan/v2ray-core/common/cmdarg"
	"github.com/chengshanshan/v2ray-core/main/commands/base"
)

var CmdRun = &base.Command{
	CustomFlags: true,
	UsageLine:   "{{.Exec}} run [-c config.json] [-d dir]",
	Short:       "run V2Ray with config",
	Long: `
Run V2Ray with config.
{{.Exec}} will also use the config directory specified by enviroment
variable "v2ray.location.confdir". If no config found,it tries
to load config from one of below:

	1. The default "config.json" in the current directory
	2. The config file from ENV "v2ray.location.config"
	3. The stdin if all failed above

Argument:

	-c, -config <file>
		Config file for V2Ray. Multiple assign is accepted.

	-d, -confdir <dir>
		A directory with config files. Multiple assign is accepted.
	
	-r
		Load confdir recursively.
	
	-format <format>
		Format of config input. (default "auto)

Examples:

	{{.Exec}} {{.LongName}} -c config.json
	{{.Exec}} {{.LongName}} -d path/to/dir

Use "{{.Exec}} help format-loader" for more information about format.
 	`,
	Run: executeRun,
}

var (
	configFiles          cmdarg.Arg
	configDirs           cmdarg.Arg
	configFormat         *string
	configDirRecursively *bool
)

func setConfigFlags(cmd *base.Command) {
	configFormat = cmd.Flag.String("format", core.FormatAuto, "")
	configDirRecursively = cmd.Flag.Bool("r", false, "")

	cmd.Flag.Var(&configFiles, "config", "")
	cmd.Flag.Var(&configFiles, "c", "")
	cmd.Flag.Var(&configDirs, "confdir", "")
	cmd.Flag.Var(&configDirs, "d", "")
}

func executeRun(cmd *base.Command, args []string) {
	setConfigFlags(cmd)

}
