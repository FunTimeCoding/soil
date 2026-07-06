package sweep

import (
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"os"
	"path/filepath"
)

func DeleteSource(sessionIdentifier string) {
	base := sourcePath()
	entries, e := os.ReadDir(base)

	if e != nil {
		return
	}

	target := join.Empty(sessionIdentifier, constant.NotationLogExtension)

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		path := filepath.Join(base, entry.Name(), target)

		if e := os.Remove(path); e != nil {
			continue
		}

		d := filepath.Join(base, entry.Name(), sessionIdentifier)

		if e := os.RemoveAll(d); e != nil {
			continue
		}
	}
}
