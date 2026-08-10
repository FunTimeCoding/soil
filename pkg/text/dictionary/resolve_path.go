package dictionary

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/system"
	systemConstant "github.com/funtimecoding/soil/pkg/system/constant"
	text "github.com/funtimecoding/soil/pkg/text/constant"
	"path/filepath"
)

func ResolvePath() string {
	candidates := []string{
		text.DictionaryFile,
		filepath.Join(systemConstant.DocumentPath, text.DictionaryFile),
	}
	result := system.FirstFile(candidates...)

	if result == "" {
		panic(fmt.Sprintf("dictionary not found: %v", candidates))
	}

	return result
}
