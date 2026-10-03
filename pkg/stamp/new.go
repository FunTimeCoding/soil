package stamp

import (
	"github.com/funtimecoding/soil/pkg/stamp/constant"
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
)

func New(
	version string,
	gitHash string,
	date string,
) *Stamp {
	b := fromBuildInformation()

	return &Stamp{
		Version:   firstSet(version, b.Version, constant.DefaultVersion),
		GitHash:   firstSet(gitHash, b.GitHash, constant.DefaultGitHash),
		BuildDate: firstSet(date, b.BuildDate, constant.DefaultDate),
		Module:    firstSet(constant.Module, b.Module),
		Dirty:     constant.Dirty == stringsConstant.BooleanTrue || b.Dirty,
	}
}
