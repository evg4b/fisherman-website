package handlers_test

import (
	"context"
	"errors"
	"fisherman/clicontext"
	"fisherman/config"
	"fisherman/handlers"
	"fisherman/infrastructure"
	inf_mock "fisherman/mocks/infrastructure"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrePushHandler(t *testing.T) {
	fakeRepository := inf_mock.Repository{}
	fakeRepository.On("GetCurrentBranch").Return("develop", nil)
	fakeRepository.On("GetLastTag").Return("0.0.0", nil)
	fakeRepository.On("GetUser").Return(infrastructure.User{}, nil)

	fakeShell := inf_mock.Shell{}

	assert.NotPanics(t, func() {
		err := handlers.PrePushHandler(clicontext.NewContext(context.TODO(), clicontext.Args{
			GlobalVariables: map[string]interface{}{},
			Config: &config.FishermanConfig{
				Hooks: config.HooksConfig{},
			},
			Repository: &fakeRepository,
			Shell:      &fakeShell,
			App:        &clicontext.AppInfo{},
		}), []string{})
		assert.NoError(t, err)
	})
}

func TestPrePushHandler_VariablesError(t *testing.T) {
	fakeRepository := inf_mock.Repository{}
	fakeRepository.On("GetCurrentBranch").Return("develop", nil)
	fakeRepository.On("GetLastTag").Return("0.0.0", errors.New("fail"))
	fakeRepository.On("GetUser").Return(infrastructure.User{}, nil)

	fakeShell := inf_mock.Shell{}

	assert.NotPanics(t, func() {
		err := handlers.PrePushHandler(clicontext.NewContext(context.TODO(), clicontext.Args{
			GlobalVariables: map[string]interface{}{},
			Config: &config.FishermanConfig{
				Hooks: config.HooksConfig{},
			},
			Repository: &fakeRepository,
			Shell:      &fakeShell,
			App:        &clicontext.AppInfo{},
		}), []string{})
		assert.Error(t, err, "fail")
	})
}
