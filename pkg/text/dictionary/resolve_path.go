package dictionary

import (
	"fmt"
	system "github.com/funtimecoding/soil/pkg/system/constant"
	"github.com/funtimecoding/soil/pkg/text/dictionary/constant"
	"os"
	"path/filepath"
)

func ResolvePath() string {
	candidates := []string{
		constant.File,
		filepath.Join(system.DocumentPath, constant.File),
	}

	for _, p := range candidates {
		if _, e := os.Stat(p); e == nil {
			return p
		}
	}

	panic(fmt.Sprintf("dictionary not found: %v", candidates))
}
