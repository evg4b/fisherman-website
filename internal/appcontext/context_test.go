package appcontext_test

import (
	"bytes"
	"context"
	"fisherman/internal"
	. "fisherman/internal/appcontext"
	"fisherman/internal/constants"
	"fisherman/pkg/vcs"
	"fisherman/testing/mocks"
	"fisherman/testing/testutils"
	"fmt"
	"runtime"
	"testing"

	"github.com/go-errors/errors"
	"github.com/stretchr/testify/assert"
)

func TestContext_Files(t *testing.T) {
	expectedFs := mocks.NewFilesystemMock(t)
	ctx := NewContext(
		WithFileSystem(expectedFs),
		WithRepository(mocks.NewRepositoryMock(t)),
	)

	actualFs := ctx.Files()

	assert.Equal(t, expectedFs, actualFs)
}

func TestContext_Repository(t *testing.T) {
	expectedRepo := mocks.NewRepositoryMock(t)
	ctx := NewContext(
		WithFileSystem(mocks.NewFilesystemMock(t)),
		WithRepository(expectedRepo),
	)

	actualRepo := ctx.Repository()

	assert.Equal(t, expectedRepo, actualRepo)
}

func TestContext_Args(t *testing.T) {
	expectedArgs := []string{"param"}
	ctx := NewContext(
		WithFileSystem(mocks.NewFilesystemMock(t)),
		WithRepository(mocks.NewRepositoryMock(t)),
		WithArgs(expectedArgs),
	)

	actualArgs := ctx.Args()

	assert.Equal(t, expectedArgs, actualArgs)
}

func TestContext_Cwd(t *testing.T) {
	expectedCwd := "/usr/root/home"

	ctx := NewContext(
		WithFileSystem(mocks.NewFilesystemMock(t)),
		WithRepository(mocks.NewRepositoryMock(t)),
		WithCwd(expectedCwd),
	)

	actualCwd := ctx.Cwd()

	assert.Equal(t, expectedCwd, actualCwd)
}

func TestContext_Arg(t *testing.T) {
	ctx := NewContext(
		WithFileSystem(mocks.NewFilesystemMock(t)),
		WithRepository(mocks.NewRepositoryMock(t)),
		WithArgs([]string{"fisherman", "handle", "--hook", "commit-msg", "/user/home/MESSAGE"}),
	)

	tests := []struct {
		name        string
		index       int
		expected    string
		expectedErr string
	}{
		{name: "first argument", index: 0, expected: "fisherman"},
		{name: "negative argument index", index: -1, expectedErr: "incorrect argument index"},
		{name: "second argument", index: 1, expected: "handle"},
		{name: "last argument", index: 4, expected: "/user/home/MESSAGE"},
		{name: "out of rage argument index", index: 50, expectedErr: "argument at index 50 is not provided"},
		{name: "at index 3", index: 3, expected: "commit-msg"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			value, err := ctx.Arg(tt.index)

			assert.Equal(t, tt.expected, value)
			testutils.AssertError(t, tt.expectedErr, err)
		})
	}
}

func TestContext_Output(t *testing.T) {
	expectedString := ""

	buffer := bytes.NewBufferString("")
	ctx := NewContext(
		WithFileSystem(mocks.NewFilesystemMock(t)),
		WithRepository(mocks.NewRepositoryMock(t)),
		WithOutput(buffer),
	)

	actualOutput := ctx.Output()

	fmt.Fprintln(actualOutput, expectedString)

	assert.NoError(t, actualOutput.Close())
	assert.Equal(t, expectedString, buffer.String())
}

func TestContext_Message(t *testing.T) {
	tests := []struct {
		name        string
		files       map[string]string
		expected    string
		expectedErr string
		args        []string
	}{
		{
			name:        "return message from file",
			files:       map[string]string{"filepath": "expectedMessage"},
			expected:    "expectedMessage",
			expectedErr: "",
			args:        []string{"handle", "--hook", "commit-msg", "filepath"},
		},
		{
			name:        "return message from file2",
			files:       map[string]string{},
			expected:    "",
			expectedErr: "argument at index 3 is not provided",
			args:        []string{"handle", "--hook", "commit-msg"},
		},
		{
			name:        "return message from file",
			files:       map[string]string{},
			expected:    "",
			expectedErr: "file does not exist",
			args:        []string{"handle", "--hook", "commit-msg", "filepath"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := NewContext(
				WithFileSystem(testutils.FsFromMap(t, tt.files)),
				WithRepository(mocks.NewRepositoryMock(t)),
				WithArgs(tt.args),
			)

			actual, err := ctx.Message()

			assert.Equal(t, tt.expected, actual)
			testutils.AssertError(t, tt.expectedErr, err)
		})
	}
}

func TestContext_Stop(t *testing.T) {
	ctx := NewContext(
		WithFileSystem(mocks.NewFilesystemMock(t)),
		WithRepository(mocks.NewRepositoryMock(t)),
	)

	ctx.Cancel()

	assert.Equal(t, context.Canceled, ctx.Err())
}

func TestContext_Value(t *testing.T) {
	key := "this-is-key"
	expected := "this-is-value"

	ctx := NewContext(
		WithBaseContext(context.WithValue(context.Background(), key, expected)), //nolint
		WithFileSystem(mocks.NewFilesystemMock(t)),
		WithRepository(mocks.NewRepositoryMock(t)),
	)

	data := ctx.Value(key)

	assert.Equal(t, expected, data)
}

func TestContext_Deadline(t *testing.T) {
	ctx := NewContext(
		WithBaseContext(context.Background()),
		WithFileSystem(mocks.NewFilesystemMock(t)),
		WithRepository(mocks.NewRepositoryMock(t)),
	)

	data, ok := ctx.Deadline()

	assert.NotNil(t, data)
	assert.False(t, ok)
}

func TestContext_Done(t *testing.T) {
	ctx := NewContext(
		WithBaseContext(context.Background()),
		WithFileSystem(mocks.NewFilesystemMock(t)),
		WithRepository(mocks.NewRepositoryMock(t)),
	)

	chanell := ctx.Done()

	assert.NotNil(t, chanell)
}

func TestContext_GlobalVariables(t *testing.T) {
	tests := []struct {
		name        string
		expected    map[string]interface{}
		repository  internal.Repository
		expectedErr string
	}{
		{
			name: "GetLastTag returns error",
			repository: mocks.NewRepositoryMock(t).
				GetLastTagMock.Return("", errors.New("GetLastTag error")),
			expected:    nil,
			expectedErr: "GetLastTag error",
		},
		{
			name: "GetCurrentBranch returns error",
			repository: mocks.NewRepositoryMock(t).
				GetLastTagMock.Return("1.0.0", nil).
				GetCurrentBranchMock.Return("", errors.New("GetCurrentBranch error")),
			expected:    nil,
			expectedErr: "GetCurrentBranch error",
		},
		{
			name: "GetUser returns error",
			repository: mocks.NewRepositoryMock(t).
				GetLastTagMock.Return("1.0.0", nil).
				GetCurrentBranchMock.Return("refs/head/develop", nil).
				GetUserMock.Return(vcs.User{}, errors.New("GetUser error")),
			expected:    nil,
			expectedErr: "GetUser error",
		},
		{
			name: "GetUser returns error",
			repository: mocks.NewRepositoryMock(t).
				GetLastTagMock.Return("1.0.0", nil).
				GetCurrentBranchMock.Return("refs/head/develop", nil).
				GetUserMock.Return(vcs.User{UserName: "evg4b", Email: "evg4b@mail.com"}, nil),
			expected: map[string]interface{}{
				constants.UserEmailVariable:        "evg4b@mail.com",
				constants.UserNameVariable:         "evg4b",
				constants.FishermanVersionVariable: constants.Version,
				constants.CwdVariable:              "~/project",
				constants.BranchNameVariable:       "refs/head/develop",
				constants.TagVariable:              "1.0.0",
				constants.OsVariable:               runtime.GOOS,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := NewContext(
				WithFileSystem(mocks.NewFilesystemMock(t)),
				WithRepository(tt.repository),
				WithCwd("~/project"),
			)

			actual, err := ctx.GlobalVariables()

			assert.Equal(t, tt.expected, actual)
			testutils.AssertError(t, tt.expectedErr, err)
		})
	}
}

func TestApplicationContext_Envs(t *testing.T) {
	expectedEnvs := []string{"VALUE1=123", "VALUE4=234234"}
	ctx := NewContext(
		WithFileSystem(mocks.NewFilesystemMock(t)),
		WithRepository(mocks.NewRepositoryMock(t)),
		WithEnv(expectedEnvs),
	)

	actualEnvs := ctx.Env()

	assert.Equal(t, expectedEnvs, actualEnvs)
}
