package sweep

import (
	"github.com/funtimecoding/soil/pkg/system"
	"path/filepath"
)

func sourcePath() string {
	return filepath.Join(
		system.Home(),
		".claude",
		"projects",
	)
}
