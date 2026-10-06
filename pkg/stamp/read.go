package stamp

import (
	gitConstant "github.com/funtimecoding/soil/pkg/git/constant"
	"github.com/funtimecoding/soil/pkg/stamp/constant"
	"runtime/debug"
)

func Read(i *debug.BuildInfo) *Stamp {
	result := &Stamp{
		Version:    constant.DefaultVersion,
		GitHash:    constant.DefaultGitHash,
		CommitDate: constant.DefaultDate,
	}

	if i == nil {
		return result
	}

	result.Module = i.Main.Path

	if i.Main.Version != "" && i.Main.Version != constant.DevelopmentVersion {
		result.Version = i.Main.Version
	}

	for _, s := range i.Settings {
		switch s.Key {
		case constant.RevisionKey:
			if len(s.Value) >= gitConstant.HashLength {
				result.GitHash = s.Value[:gitConstant.HashLength]
			}
		case constant.TimeKey:
			result.CommitDate = s.Value
		case constant.ModifiedKey:
			result.Dirty = s.Value == constant.ModifiedValue
		}
	}

	if result.GitHash == constant.DefaultGitHash {
		result.GitHash = pseudoVersionHash(result.Version)
	}

	return result
}
