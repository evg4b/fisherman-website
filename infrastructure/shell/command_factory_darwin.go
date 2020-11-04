package shell

import (
	"os/exec"
	"strings"
)

const LineBreak = "\n"
const PathVariableSeparator = ":"

func CommandFactory(commands []string) (*exec.Cmd, error) {
	ps, err := exec.LookPath("bash")
	if err != nil {
		return nil, err
	}

	command := strings.Join(commands, LineBreak)

	return exec.Command(ps, "-c", command), nil
}
