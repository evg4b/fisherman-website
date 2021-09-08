package configuration_test

import (
	. "fisherman/internal/configuration"
	"fisherman/internal/constants"
	"fisherman/internal/rules"
	"fisherman/pkg/log"
	"fisherman/testing/mocks"
	"fisherman/testing/testutils"
	"fmt"
	"os/user"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfigLoader_FindConfigFiles(t *testing.T) {
	usr := user.User{HomeDir: filepath.Join("/", "usr", "home")}
	cwd := filepath.Join("/", "usr", "home", "documents", "repo")

	localConfig := GetConfigFolder(&usr, cwd, LocalMode)
	repoConfig := GetConfigFolder(&usr, cwd, RepoMode)
	globalConfig := GetConfigFolder(&usr, cwd, GlobalMode)

	tests := []struct {
		name        string
		files       []string
		expected    map[string]string
		expectedErr string
	}{
		{
			name: "mere then one config file",
			files: []string{
				filepath.Join(localConfig, constants.AppConfigNames[0]),
				filepath.Join(localConfig, constants.AppConfigNames[1]),
				filepath.Join(repoConfig, constants.AppConfigNames[0]),
				filepath.Join(globalConfig, constants.AppConfigNames[0]),
			},
			expectedErr: fmt.Sprintf("more then one config file specifies in folder '%s'", localConfig),
		},
		{
			name: "correct files loading",
			files: []string{
				filepath.Join(localConfig, constants.AppConfigNames[0]),
				filepath.Join(repoConfig, constants.AppConfigNames[0]),
				filepath.Join(globalConfig, constants.AppConfigNames[0]),
			},
			expected: map[string]string{
				LocalMode:  filepath.Join(localConfig, constants.AppConfigNames[0]),
				RepoMode:   filepath.Join(repoConfig, constants.AppConfigNames[0]),
				GlobalMode: filepath.Join(globalConfig, constants.AppConfigNames[0]),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var fs = testutils.FsFromSlice(t, tt.files)
			loaded := NewLoader(&usr, cwd, fs)

			actual, err := loaded.FindConfigFiles()

			testutils.CheckError(t, tt.expectedErr, err)
			assert.EqualValues(t, tt.expected, actual)
		})
	}
}

func TestGetConfigFolder(t *testing.T) {
	usr := user.User{HomeDir: filepath.Join("/", "usr", "home")}
	cwd := filepath.Join("/", "usr", "home", "documents", "repo")

	tests := []struct {
		name        string
		usr         *user.User
		cwd         string
		mode        string
		expected    string
		shouldPanic bool
	}{
		{
			name:     "local mode",
			usr:      &usr,
			cwd:      cwd,
			mode:     LocalMode,
			expected: filepath.Join(cwd, ".git"),
		},
		{
			name:     "global mode",
			usr:      &usr,
			cwd:      cwd,
			mode:     GlobalMode,
			expected: usr.HomeDir,
		},
		{
			name:     "repository mode",
			usr:      &usr,
			cwd:      cwd,
			mode:     RepoMode,
			expected: cwd,
		},
		{
			name:        "unknown mode",
			usr:         &usr,
			cwd:         cwd,
			mode:        "unknown mode",
			shouldPanic: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.shouldPanic {
				assert.Panics(t, func() {
					_ = GetConfigFolder(tt.usr, tt.cwd, tt.mode)
				})
			} else {
				actual := GetConfigFolder(tt.usr, tt.cwd, tt.mode)

				assert.Equal(t, tt.expected, actual)
			}
		})
	}
}

func TestConfigLoader_Load(t *testing.T) {
	usr := user.User{HomeDir: filepath.Join("/", "usr", "home")}
	cwd := filepath.Join("/", "usr", "home", "documents", "repo")

	config := `
variables:
  name: value

hooks:
  pre-push:
    rules:
      - type: shell-script
        name: Demo
        commands:
         - echo '1213123' >> log.txt
         - exit 1`

	fs := testutils.FsFromMap(t, map[string]string{
		"GlobalConfig":      config,
		"GlobalConfigError": "asd['",
	})

	tests := []struct {
		name        string
		loader      *ConfigLoader
		files       map[string]string
		expected    *FishermanConfig
		expectedErr string
	}{
		{
			name:   "",
			loader: NewLoader(&usr, cwd, mocks.NewFilesystemMock(t)),
			files:  map[string]string{},
			expected: &FishermanConfig{
				Output: log.DefaultOutputConfig,
			},
		},
		{
			name:   "file reader error",
			loader: NewLoader(&usr, cwd, fs),
			files: map[string]string{
				GlobalMode: "GlobalConfig3",
			},
			expectedErr: "open GlobalConfig3: file does not exist",
		},
		{
			name:   "correct decoding",
			loader: NewLoader(&usr, cwd, fs),
			files: map[string]string{
				GlobalMode: "GlobalConfig",
			},
			expected: &FishermanConfig{
				Output: log.DefaultOutputConfig,
				GlobalVariables: map[string]interface{}{
					"name": "value",
				},
				Hooks: HooksConfig{
					PrePushHook: &HookConfig{
						Rules: []Rule{
							&rules.ShellScript{
								BaseRule: rules.BaseRule{
									Type: "shell-script",
								},
								Name: "Demo",
								Commands: []string{
									"echo '1213123' >> log.txt",
									"exit 1",
								},
							},
						},
					},
				},
			},
		},
		{
			name:   "decoding error",
			loader: NewLoader(&usr, cwd, fs),
			files: map[string]string{
				GlobalMode: "GlobalConfigError",
			},
			expectedErr: "yaml: unmarshal errors:\n  line 1: cannot unmarshal !!str `asd['` into configuration.FishermanConfig",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := tt.loader.Load(tt.files)

			testutils.CheckError(t, tt.expectedErr, err)
			if tt.expectedErr == "" {
				assert.Equal(t, tt.expected, actual)
			}
		})
	}
}

func TestConfigLoader_Load_Correct_Merging(t *testing.T) {
	usr := user.User{HomeDir: filepath.Join("/", "usr", "home")}
	cwd := filepath.Join("/", "usr", "home", "documents", "repo")

	fs := testutils.FsFromMap(t, map[string]string{
		"global.yaml": `
variables:
  var1: global
  var2: global
  var3: global
  var4: global`,
		"repo.yaml": `
variables:
  var1: repo
  var2: repo`,
		"local.yaml": `
variables:
  var1: local
  var3: local`,
	})

	loader := NewLoader(&usr, cwd, fs)

	files := map[string]string{
		GlobalMode: "global.yaml",
		LocalMode:  "local.yaml",
		RepoMode:   "repo.yaml",
	}

	actual, err := loader.Load(files)

	assert.NoError(t, err)
	assert.Equal(t, &FishermanConfig{
		Output: log.DefaultOutputConfig,
		GlobalVariables: map[string]interface{}{
			"var1": "local",
			"var2": "repo",
			"var3": "local",
			"var4": "global",
		},
	}, actual)
}
