package shell

import (
	"context"
	"os/exec"

	"golang.org/x/text/encoding"
)

type BashStrategy struct{}

const bashBin = "bash"

func Bash() *BashStrategy {
	return &BashStrategy{}
}

func (s *BashStrategy) GetName() string {
	return bashBin
}

func (s *BashStrategy) GetCommand(ctx context.Context) *exec.Cmd {
	return exec.CommandContext(ctx, bashBin)
}

func (s *BashStrategy) ArgsWrapper(args []string) []string {
	return args
}

func (s *BashStrategy) EnvWrapper(env []string) []string {
	return env
}

func (s *BashStrategy) GetEncoding() encoding.Encoding {
	return encoding.Nop
}
