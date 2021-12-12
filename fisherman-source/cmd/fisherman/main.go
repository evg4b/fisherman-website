package main

import (
	"context"
	"fisherman/internal"
	"fisherman/internal/app"
	"fisherman/internal/commands/handle"
	"fisherman/internal/commands/initialize"
	"fisherman/internal/commands/remove"
	"fisherman/internal/commands/version"
	"fisherman/internal/configuration"
	"fisherman/internal/expression"
	"fisherman/internal/utils"
	"fisherman/pkg/guards"
	"fisherman/pkg/log"
	"fisherman/pkg/vcs"
	"os"
	"os/user"
	"runtime"

	"github.com/go-git/go-billy/v5/osfs"
)

const fatalExitCode = 1

func main() {
	defer utils.PanicInterceptor(func(recovered interface{}) {
		log.Errorf("Fatal error: %s", recovered)
		if err, ok := recovered.(error); ok {
			log.DumpError(err)
		}

		os.Exit(fatalExitCode)
	})

	usr, err := user.Current()
	guards.NoError(err)

	cwd, err := os.Getwd()
	guards.NoError(err)

	executablePath, err := os.Executable()
	guards.NoError(err)

	fs := osfs.New("")

	configs, err := configuration.FindConfigFiles(usr, cwd, fs)
	guards.NoError(err)

	config, err := configuration.Load(fs, configs)
	guards.NoError(err)

	log.Configure(config.Output)

	ctx := context.Background()
	engine := expression.NewGoExpressionEngine()

	appInfo := internal.AppInfo{
		Executable: executablePath,
		Cwd:        cwd,
		Configs:    configs,
	}

	repo := vcs.NewRepository(vcs.WithFsRepo(cwd))

	fishermanApp := app.NewFishermanApp(
		app.WithCommands([]internal.CliCommand{
			initialize.NewCommand(fs, appInfo, usr),
			handle.NewCommand(
				handle.WithExpressionEngine(engine),
				handle.WithHooksConfig(&config.Hooks),
				handle.WithGlobalVars(map[string]interface{}{
					// TODO: provide variables
				}),
				handle.WithCwd(cwd),
				handle.WithFileSystem(fs),
				handle.WithRepository(repo),
				handle.WithArgs(os.Args[1:]),
				handle.WithEnv(os.Environ()),
				handle.WithWorkersCount(uint(runtime.NumCPU())),
				handle.WithConfigFiles(configs),
				handle.WithOutput(os.Stdout),
			),
			remove.NewCommand(fs, appInfo, usr),
			version.NewCommand(),
		}),
		app.WithFs(fs),
		app.WithOutput(os.Stdout),
		app.WithRepository(repo),
		app.WithEnv(os.Environ()),
		app.WithSistemInterruptSignals(),
	)

	if err = fishermanApp.Run(ctx, os.Args[1:]); err != nil {
		panic(err)
	}
}
