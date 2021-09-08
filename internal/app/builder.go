package app

import (
	i "fisherman/internal"
	"fisherman/pkg/guards"
	"io"

	"github.com/go-git/go-billy/v5"
)

type Builder struct {
	fs       billy.Filesystem
	shell    i.Shell
	repo     i.Repository
	output   io.Writer
	commands CliCommands
	cwd      string
}

func NewAppBuilder() *Builder {
	return &Builder{
		output: io.Discard,
	}
}

func (rb *Builder) WithCommands(commands []i.CliCommand) *Builder {
	rb.commands = commands

	return rb
}

func (rb *Builder) WithCwd(cwd string) *Builder {
	rb.cwd = cwd

	return rb
}

func (rb *Builder) WithFs(fs billy.Filesystem) *Builder {
	rb.fs = fs

	return rb
}

func (rb *Builder) WithOutput(output io.Writer) *Builder {
	rb.output = output

	return rb
}

func (rb *Builder) WithShell(shell i.Shell) *Builder {
	rb.shell = shell

	return rb
}

func (rb *Builder) WithRepository(repo i.Repository) *Builder {
	rb.repo = repo

	return rb
}

func (rb *Builder) Build() FishermanApp {
	guards.ShouldBeDefined(rb.fs, "FileSystem should be configured")
	guards.ShouldBeDefined(rb.shell, "Shell should be configured")
	guards.ShouldBeDefined(rb.repo, "Repository should be configured")
	guards.ShouldBeDefined(rb.commands, "Commands should be configured")

	return FishermanApp{
		fs:       rb.fs,
		shell:    rb.shell,
		repo:     rb.repo,
		output:   rb.output,
		commands: rb.commands,
	}
}
