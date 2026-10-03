package stamp

import (
	gitConstant "github.com/funtimecoding/soil/pkg/git/constant"
	"github.com/funtimecoding/soil/pkg/stamp/constant"
	"runtime/debug"
)

func fromBuildInformation() *Stamp {
	result := &Stamp{}
	i, okay := debug.ReadBuildInfo()

	if !okay {
		return result
	}

	result.Module = i.Main.Path

	if i.Main.Version != constant.DevelopmentVersion {
		result.Version = i.Main.Version
	}

	for _, s := range i.Settings {
		switch s.Key {
		case constant.RevisionKey:
			if len(s.Value) >= gitConstant.HashLength {
				result.GitHash = s.Value[:gitConstant.HashLength]
			}
		case constant.TimeKey:
			result.BuildDate = s.Value
		case constant.ModifiedKey:
			result.Dirty = s.Value == constant.ModifiedValue
		}
	}

	return result
}
