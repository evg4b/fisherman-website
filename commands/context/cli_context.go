package context

import (
	"fisherman/config"
	"fisherman/infrastructure/git"
	"fisherman/infrastructure/io"
	"fisherman/infrastructure/logger"
	"os/user"
)

// CommandContext is cli context structure
type CommandContext struct {
	repoInfo         *git.RepositoryInfo
	usr              *user.User
	cwd              string
	config           *config.FishermanConfig
	appPath          string
	globalConfigPath string
	repoConfigPath   string
	localConfigPath  string
	path             string
	FileAccessor     io.FileAccessor
	Logger           logger.Logger
}

// CliCommandContextParams is structure for params in cli command context constructor
type CliCommandContextParams struct {
	RepoInfo     *git.RepositoryInfo
	FileAccessor io.FileAccessor
	Usr          *user.User
	Cwd          string
	AppPath      string
	ConfigInfo   *config.LoadInfo
	Path         string
	Logger       logger.Logger
}

// NewContext constructor for cli command context
func NewContext(params CliCommandContextParams) *CommandContext {
	configInfo := params.ConfigInfo
	return &CommandContext{
		params.RepoInfo,
		params.Usr,
		params.Cwd,
		configInfo.Config,
		params.AppPath,
		configInfo.GlobalConfigPath,
		configInfo.RepoConfigPath,
		configInfo.LocalConfigPath,
		params.Path,
		params.FileAccessor,
		params.Logger,
	}
}
