package dictionary

import (
	"github.com/funtimecoding/soil/pkg/text/dictionary/constant"
	"path/filepath"
	"strings"
)

func IncludeFile(name string) bool {
	e := filepath.Ext(name)

	if e == "" {
		return constant.NoExtension[name]
	}

	for k := range constant.Prefix {
		if strings.HasPrefix(name, k) {
			return true
		}
	}

	return constant.Extension[strings.ToLower(e)]
}
