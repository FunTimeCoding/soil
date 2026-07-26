package dictionary

import (
	"fmt"
	system "github.com/funtimecoding/soil/pkg/system/constant"
	text "github.com/funtimecoding/soil/pkg/text/constant"
	"os"
	"path/filepath"
)

func ResolvePath() string {
	candidates := []string{
		text.DictionaryFile,
		filepath.Join(system.DocumentPath, text.DictionaryFile),
	}

	for _, p := range candidates {
		if _, e := os.Stat(p); e == nil {
			return p
		}
	}

	panic(fmt.Sprintf("dictionary not found: %v", candidates))
}
