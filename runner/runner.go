package runner

import (
	"fisherman/commands"
	"fisherman/commands/handle"
	initc "fisherman/commands/init"
	"fisherman/config"
	"fisherman/constants"
	"fisherman/infrastructure/git"
	"fisherman/infrastructure/io"
	"flag"
	"fmt"
	"os"
	"os/user"
)

type Runner struct {
	fileAccessor io.FileAccessor
	usr          *user.User
}

func NewRunner(fileAccessor io.FileAccessor, usr *user.User) *Runner {
	return &Runner{fileAccessor, usr}
}

func (runner *Runner) Run(args []string) error {
	commandList := registerCommands()
	if len(args) < 2 {
		fmt.Print(constants.Logo)
		flag.Parse()
		flag.PrintDefaults()
		return nil
	}
	return runner.runInternal(args[1:], args[0], commandList)
}

func (runner *Runner) runInternal(args []string, appPath string, commandList []commands.CliCommand) error {
	subcommand := args[0]
	for _, command := range commandList {
		if command.Name() == subcommand {
			if err := command.Init(args[1:]); err != nil {
				return err
			}
			context, err := runner.buildContext(appPath)
			if err != nil {
				return err
			}
			return command.Run(context)
		}
	}
	return fmt.Errorf("unknown subcommand: %s", subcommand)
}

func (runner *Runner) buildContext(appPath string) (*commands.CliCommandContext, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	info, err := git.GetRepositoryInfo(cwd)
	if err != nil {
		return nil, err
	}
	configInfo, err := config.LoadConfig(cwd, runner.usr, runner.fileAccessor)
	if err != nil {
		return nil, err
	}
	context := commands.NewContext(commands.CliCommandContextParams{
		RepoInfo:     info,
		FileAccessor: runner.fileAccessor,
		Usr:          runner.usr,
		Cwd:          cwd,
		AppPath:      appPath,
		ConfigInfo:   configInfo,
	})
	return context, nil
}

func registerCommands() []commands.CliCommand {
	errorHandlingMode := flag.ExitOnError
	return []commands.CliCommand{
		initc.NewCommand(errorHandlingMode),
		handle.NewCommand(errorHandlingMode),
	}
}
