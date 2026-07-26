package dictionary

import (
	"github.com/funtimecoding/soil/pkg/text/constant"
	"path/filepath"
	"strings"
)

func IncludeFile(name string) bool {
	e := filepath.Ext(name)

	if e == "" {
		return constant.DictionaryNoExtension[name]
	}

	for k := range constant.DictionaryPrefix {
		if strings.HasPrefix(name, k) {
			return true
		}
	}

	return constant.DictionaryExtension[strings.ToLower(e)]
}
