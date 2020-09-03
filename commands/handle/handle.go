package handle

import (
	c "fisherman/constants"
	"fisherman/handlers"
	handler "fisherman/handlers"
	"fisherman/infrastructure/io"
	"fisherman/infrastructure/reporter"
	"flag"
)

// Command is structure for storage information about handle command
type Command struct {
	fs       *flag.FlagSet
	hook     string
	args     []string
	reporter reporter.Reporter
	handlers map[string]handler.HookHandler
}

// NewCommand is constructor for handle command
func NewCommand(handling flag.ErrorHandling, r reporter.Reporter, f io.FileAccessor) *Command {
	fs := flag.NewFlagSet("handle", handling)
	c := &Command{
		fs:       fs,
		reporter: r,
		handlers: map[string]handler.HookHandler{
			c.ApplyPatchMsgHook:     handlers.ApplyPatchMsgHandler,
			c.CommitMsgHook:         handlers.CommitMsgHandler,
			c.FsMonitorWatchmanHook: handlers.FsMonitorWatchmanHandler,
			c.PostUpdateHook:        handlers.PostUpdateHandler,
			c.PreApplyPatchHook:     handlers.PreApplyPatchHandler,
			c.PreCommitHook:         handlers.PreCommitHandler,
			c.PrePushHook:           handlers.PrePushHandler,
			c.PreRebaseHook:         handlers.PreRebaseHandler,
			c.PreReceiveHook:        handlers.PreReceiveHandler,
			c.PrepareCommitMsgHook:  handlers.PrepareCommitMsgHandler,
			c.UpdateHook:            handlers.UpdateHandler,
		},
	}
	fs.StringVar(&c.hook, "hook", "", "")
	return c
}

// Name returns command name
func (c *Command) Name() string {
	return c.fs.Name()
}
