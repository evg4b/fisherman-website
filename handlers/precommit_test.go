package handlers_test

import (
	"fisherman/commands"
	"fisherman/constants"
	"fisherman/handlers"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPreCommitHandler(t *testing.T) {
	constants.Version = "1.0.1"
	assert.NotPanics(t, func() {
		err := handlers.PreCommitHandler(&commands.CommandContext{}, []string{})
		assert.Error(t, err, "This hook is not supported in version 1.0.1.")
	})
}