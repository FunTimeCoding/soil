package installed

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	stampConstant "github.com/funtimecoding/soil/pkg/stamp/constant"
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"strings"
)

func (b *Binary) Link(flags string) {
	for _, field := range strings.Fields(strings.Trim(flags, `"`)) {
		key, value, found := strings.Cut(field, stringsConstant.Equals)

		if !found {
			continue
		}

		switch key {
		case constant.MainVersionVariable:
			b.Version = value
		case constant.MainGitHashVariable:
			b.Hash = value
		case stampConstant.ModuleVariable:
			b.Module = value
		case stampConstant.DirtyVariable:
			b.Dirty = value == stringsConstant.BooleanTrue
		}
	}
}
