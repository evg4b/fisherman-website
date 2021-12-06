package app_test

import (
	"context"
	"fisherman/internal"
	. "fisherman/internal/app"
	"fisherman/testing/mocks"
	"fisherman/testing/testutils"
	"fmt"
	"io"
	"testing"

	"github.com/go-errors/errors"

	"github.com/stretchr/testify/assert"
)

func TestRunner_Run(t *testing.T) {
	tests := []struct {
		name        string
		args        []string
		commands    []internal.CliCommand
		expectedErr string
	}{
		{
			name: "should run called commnad and return its error",
			args: []string{"init"},
			commands: []internal.CliCommand{
				makeCommand(t, "handle"),
				makeCommand(t, "remove"),
				makeExpectedCommand(t, "init", errors.New("expected error")),
			},
			expectedErr: "expected error",
		},
		{
			name: "should run called commnad and return nil when command executed witout error",
			args: []string{"init"},
			commands: []internal.CliCommand{
				makeCommand(t, "handle"),
				makeCommand(t, "remove"),
				makeExpectedCommand(t, "init", nil),
			},
		},
		{
			name: "should return error when command not found",
			args: []string{"not"},
			commands: []internal.CliCommand{
				makeCommand(t, "handle"),
				makeCommand(t, "remove"),
				makeCommand(t, "init"),
			},
			expectedErr: "unknown command: not",
		},
		{
			name:        "Should return error when command not registered",
			args:        []string{"not"},
			commands:    []internal.CliCommand{},
			expectedErr: "unknown command: not",
		},
		{
			name: "should not return error when commnad not specified",
			args: []string{},
			commands: []internal.CliCommand{
				makeCommand(t, "handle"),
				makeCommand(t, "remove"),
				makeCommand(t, "init"),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			appInstance := NewFishermanApp(
				WithCommands(tt.commands),
				WithFs(mocks.NewFilesystemMock(t)),
				WithRepository(mocks.NewRepositoryMock(t)),
				WithCwd("/"),
				WithOutput(io.Discard),
			)

			assert.NotPanics(t, func() {
				err := appInstance.Run(context.TODO(), tt.args)
				testutils.AssertError(t, tt.expectedErr, err)
			})
		})
	}
}

func makeCommand(t *testing.T, name string) *mocks.CliCommandMock {
	t.Helper()

	return mocks.NewCliCommandMock(t).
		NameMock.Return(name).
		InitMock.Return(nil).
		DescriptionMock.Return(fmt.Sprintf("This is %s command", name))
}

func makeExpectedCommand(t *testing.T, name string, err error) *mocks.CliCommandMock {
	t.Helper()

	return makeCommand(t, name).
		RunMock.Return(err)
}
