package appcontext

import (
	"context"
	"fisherman/internal"
	"fisherman/pkg/guards"
	"io"

	"github.com/evg4b/linebyline"
	"github.com/go-git/go-billy/v5"
)

type ContextBuilder struct {
	cwd    string
	fs     billy.Filesystem
	shell  internal.Shell
	repo   internal.Repository
	args   []string
	output io.Writer
	ctx    context.Context
}

func NewContextBuilder() *ContextBuilder {
	return &ContextBuilder{
		output: io.Discard,
		ctx:    context.TODO(),
		args:   []string{},
	}
}

func (cb *ContextBuilder) WithFileSystem(fileSystem billy.Filesystem) *ContextBuilder {
	cb.fs = fileSystem

	return cb
}

func (cb *ContextBuilder) WithCwd(cwd string) *ContextBuilder {
	cb.cwd = cwd

	return cb
}

func (cb *ContextBuilder) WithShell(shell internal.Shell) *ContextBuilder {
	cb.shell = shell

	return cb
}

func (cb *ContextBuilder) WithRepository(repository internal.Repository) *ContextBuilder {
	cb.repo = repository

	return cb
}

func (cb *ContextBuilder) WithArgs(args []string) *ContextBuilder {
	cb.args = args

	return cb
}

func (cb *ContextBuilder) WithOutput(output io.Writer) *ContextBuilder {
	cb.output = output

	return cb
}

func (cb *ContextBuilder) WithContext(ctx context.Context) *ContextBuilder {
	cb.ctx = ctx

	return cb
}

func (cb *ContextBuilder) Build() *ApplicationContext {
	guards.ShouldBeDefined(cb.fs, "FileSystem should be connfigured")
	guards.ShouldBeDefined(cb.shell, "Shell should be connfigured")
	guards.ShouldBeDefined(cb.repo, "Repository should be connfigured")

	baseContext, cancelBaseContext := context.WithCancel(cb.ctx)

	return &ApplicationContext{
		cwd:           cb.cwd,
		fs:            cb.fs,
		shell:         cb.shell,
		repo:          cb.repo,
		args:          cb.args,
		output:        *linebyline.NewWriterGroup(cb.output),
		baseCtx:       baseContext,
		cancelBaseCtx: cancelBaseContext,
	}
}
