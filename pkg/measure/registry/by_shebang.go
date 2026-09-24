package registry

import (
	"github.com/funtimecoding/soil/pkg/measure/constant"
	"github.com/funtimecoding/soil/pkg/measure/language"
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/strings/split"
	"path/filepath"
	"slices"
	"strings"
)

func (r *Registry) ByShebang(firstLine string) *language.Language {
	if !strings.HasPrefix(firstLine, constant.ShebangPrefix) {
		return nil
	}

	fields := strings.Fields(
		strings.TrimPrefix(firstLine, constant.ShebangPrefix),
	)

	if len(fields) == 0 {
		return nil
	}

	interpreter := filepath.Base(fields[0])

	if interpreter == constant.EnvironmentInterpreter && len(fields) > 1 {
		interpreter = filepath.Base(fields[1])
	}

	interpreter = strings.TrimRight(
		split.Dot(interpreter)[0],
		stringsConstant.Digits,
	)

	for _, l := range r.languages {
		if slices.Contains(l.Shebangs, interpreter) {
			return l
		}
	}

	return nil
}
