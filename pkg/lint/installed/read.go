package installed

import (
	"debug/buildinfo"
	"github.com/funtimecoding/soil/pkg/lint/constant"
	stampConstant "github.com/funtimecoding/soil/pkg/stamp/constant"
	"path/filepath"
)

func Read(path string) (*Binary, bool) {
	i, e := buildinfo.ReadFile(path)

	if e != nil {
		return nil, false
	}

	result := New(filepath.Base(path), path, i.Path)
	result.Module = i.Main.Path
	result.Version = i.Main.Version

	for _, d := range i.Deps {
		if d.Replace != nil {
			d = d.Replace
		}

		result.Modules[d.Path] = d.Version
	}

	for _, s := range i.Settings {
		switch s.Key {
		case constant.LinkerFlagsSetting:
			result.Link(s.Value)
		case stampConstant.RevisionKey:
			result.Hash = s.Value
		case stampConstant.ModifiedKey:
			result.Dirty = s.Value == stampConstant.ModifiedValue
		}
	}

	return result, true
}
