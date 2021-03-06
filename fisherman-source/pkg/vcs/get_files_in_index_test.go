package vcs_test

import (
	"errors"
	"fisherman/pkg/guards"
	. "fisherman/pkg/vcs"
	"fisherman/testing/mocks"
	"fisherman/testing/testutils"
	"testing"

	"github.com/go-git/go-git/v5/storage"
	"github.com/stretchr/testify/assert"
)

func TestGitRepository_GetFilesInIndex(t *testing.T) {
	t.Run("no files", func(t *testing.T) {
		repo, _, fs, w := testutils.CreateRepo(t)

		testutils.MakeCommits(t, w, fs, map[string]map[string]string{
			"init commit": {"LICENSE": "MIT"},
			"test commit": {"demo": "this is test file"},
		})

		files, err := repo.GetFilesInIndex()

		assert.NoError(t, err)
		assert.Empty(t, files)
	})

	t.Run("excluded untracked files", func(t *testing.T) {
		repo, _, fs, w := testutils.CreateRepo(t)

		testutils.MakeCommits(t, w, fs, map[string]map[string]string{
			"init commit": {"LICENSE": "MIT"},
			"test commit": {"demo": "this is test file"},
		})

		testutils.MakeFiles(t, fs, map[string]string{
			"untracked": "untracked content",
		})

		files, err := repo.GetFilesInIndex()

		assert.NoError(t, err)
		assert.Empty(t, files)
	})

	t.Run("added files successfully", func(t *testing.T) {
		repo, _, fs, w := testutils.CreateRepo(t)

		testutils.MakeCommits(t, w, fs, map[string]map[string]string{
			"init commit": {"LICENSE": "MIT"},
			"test commit": {"demo": "this is test file"},
		})

		testutils.MakeFiles(t, fs, map[string]string{
			"tracked": "tracked content",
		})

		err := w.AddGlob(".")
		guards.NoError(err)

		files, err := repo.GetFilesInIndex()

		assert.NoError(t, err)
		assert.Equal(t, []string{"tracked"}, files)
	})

	t.Run("workdree error", func(t *testing.T) {
		expectedErr := errors.New("worktree error")
		gitMock := mocks.NewGoGitRepositoryMock(t).WorktreeMock.Return(nil, expectedErr)

		repo := NewRepository(WithFactoryMethod(func() (GoGitRepository, storage.Storer, error) {
			return gitMock, nil, nil
		}))

		_, err := repo.GetFilesInIndex()

		assert.EqualError(t, err, expectedErr.Error())
	})
}
