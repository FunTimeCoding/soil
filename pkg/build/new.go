package build

import "github.com/funtimecoding/soil/pkg/build/constant"

func New(
	version string,
	gitHash string,
	date string,
) *Build {
	if version == "" {
		version = constant.DefaultVersion
	}

	if gitHash == "" {
		gitHash = constant.DefaultGitHash
	}

	if date == "" {
		date = constant.DefaultDate
	}

	return &Build{Version: version, GitHash: gitHash, BuildDate: date}
}
