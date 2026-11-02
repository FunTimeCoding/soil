package gofix

import (
	"github.com/funtimecoding/soil/pkg/system"
	systemConstant "github.com/funtimecoding/soil/pkg/system/constant"
	textConstant "github.com/funtimecoding/soil/pkg/text/constant"
	"github.com/funtimecoding/soil/pkg/text/dictionary"
	"github.com/funtimecoding/soil/pkg/tool/gofix/constant"
	"path/filepath"
	"strings"
)

func knownWords(root string) map[string]bool {
	categories := dictionary.Parse(strings.NewReader(constant.Vocabulary))
	repository := system.FirstFile(
		filepath.Join(root, textConstant.DictionaryFile),
		filepath.Join(
			root,
			systemConstant.DocumentPath,
			textConstant.DictionaryFile,
		),
	)

	if repository != "" {
		categories = append(categories, dictionary.Read(repository)...)
	}

	result := make(map[string]bool)

	for _, c := range categories {
		for _, w := range c.Words {
			result[strings.ToLower(w)] = true
		}
	}

	return result
}
