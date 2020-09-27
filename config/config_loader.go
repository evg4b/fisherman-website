package config

import (
	"fisherman/constants"
	"fisherman/infrastructure"
	"fisherman/infrastructure/logger"
	"fisherman/utils"
	"fmt"
	"os/user"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

const gitDir = ".git"

// ConfigInfo is
type ConfigInfo struct {
	GlobalConfigPath string
	RepoConfigPath   string
	LocalConfigPath  string
}

// LoadConfig is demo
func LoadConfig(cwd string, usr *user.User, accessor infrastructure.FileAccessor) (*FishermanConfig, *ConfigInfo, error) {
	config := FishermanConfig{
		Output: logger.DefaultOutputConfig,
	}

	global, err := unmarshlIfExist(cwd, usr, GlobalMode, accessor, &config)
	if err != nil {
		return nil, nil, err
	}

	repo, err := unmarshlIfExist(cwd, usr, RepoMode, accessor, &config)
	if err != nil {
		return nil, nil, err
	}

	local, err := unmarshlIfExist(cwd, usr, LocalMode, accessor, &config)
	if err != nil {
		return nil, nil, err
	}

	loadInfo := &ConfigInfo{
		GlobalConfigPath: global,
		RepoConfigPath:   repo,
		LocalConfigPath:  local,
	}

	return &config, loadInfo, nil
}

func unmarshlIfExist(cwd string, usr *user.User, mode string, accessor infrastructure.FileAccessor, config *FishermanConfig) (string, error) {
	path, err := BuildFileConfigPath(cwd, usr, mode)
	utils.HandleCriticalError(err)

	if accessor.Exist(path) {
		data, err := accessor.Read(path)
		utils.HandleCriticalError(err)
		err = yaml.Unmarshal([]byte(data), config)
		utils.HandleCriticalError(err)
		return path, nil
	}

	return "", nil
}

// BuildFileConfigPath returns path to config by config mode
func BuildFileConfigPath(cwd string, usr *user.User, mode string) (string, error) {
	switch mode {
	case LocalMode:
		return filepath.Join(cwd, gitDir, constants.AppConfigName), nil
	case RepoMode:
		return filepath.Join(cwd, constants.AppConfigName), nil
	case GlobalMode:
		return filepath.Join(usr.HomeDir, constants.AppConfigName), nil
	default:
		return "", fmt.Errorf("Unknown config mode.")
	}
}
