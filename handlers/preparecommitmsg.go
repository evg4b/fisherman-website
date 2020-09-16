package handlers

import (
	"fisherman/commands/context"
	"regexp"
	"strings"
)

// PrepareCommitMsgHandler is a execute function for prepare-commit-msg hook
func PrepareCommitMsgHandler(ctx *context.CommandContext, args []string) error {
	config := ctx.GetHookConfiguration()
	info, err := ctx.GetGitInfo()
	if err != nil {
		panic(err)
	}

	c := config.PrepareCommitMsgHook
	if c != nil {
		message, isPresented := getPreparedMessage(c.Message, c.BranchRegExp, info.CurrentBranch)
		if isPresented {
			err = ctx.FileAccessor.WriteFile(args[0], message)
			if err != nil {
				panic(err)
			}
		}
	}

	return nil
}

func getPreparedMessage(message, regexpString, branch string) (string, bool) {
	if !isEmpty(message) {
		if !isEmpty(regexpString) {
			return regexp.MustCompile(regexpString).
				ReplaceAllString(branch, message), true
		}

		return message, true
	}
	return "", false
}

func isEmpty(data string) bool {
	return len(strings.TrimSpace(data)) == 0
}