package handlers

import (
	"fisherman/commands"
	"fisherman/config"
	iomock "fisherman/mocks/infrastructure"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrepareCommitMsgHandler(t *testing.T) {
	fakeRepository := iomock.Repository{}
	fakeRepository.On("GetCurrentBranch").Return("develop", nil)
	fakeRepository.On("GetLastTag").Return("0.0.0", nil)

	faceFileAccessor := iomock.FileAccessor{}
	faceFileAccessor.On("Read", ".git/MESSAGE").Return("[fisherman] test commit", nil)

	tests := []struct {
		name string
		args []string
		err  error
	}{
		{name: "base test", args: []string{}, err: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := commands.NewContext(commands.CliCommandContextParams{
				Config:       &config.DefaultConfig,
				Repository:   &fakeRepository,
				FileAccessor: &faceFileAccessor,
			})
			err := PrepareCommitMsgHandler(ctx, tt.args)
			assert.Equal(t, tt.err, err)
		})
	}
}
