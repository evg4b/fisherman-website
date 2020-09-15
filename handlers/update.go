package handlers

import (
	"fisherman/commands/context"
	"fisherman/constants"
	"fmt"
)

// UpdateHandler is a handler for update hook
func UpdateHandler(ctx *context.CommandContext, args []string) error {
	return fmt.Errorf("This hook is not supported in version %s.", constants.Version)
}
