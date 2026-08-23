package stamp

import "github.com/funtimecoding/soil/pkg/stamp/constant"

func New(
	version string,
	gitHash string,
	date string,
) *Stamp {
	if version == "" {
		version = constant.DefaultVersion
	}

	if gitHash == "" {
		gitHash = constant.DefaultGitHash
	}

	if date == "" {
		date = constant.DefaultDate
	}

	return &Stamp{Version: version, GitHash: gitHash, BuildDate: date}
}
