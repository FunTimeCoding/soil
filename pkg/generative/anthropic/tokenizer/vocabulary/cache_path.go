package vocabulary

import (
	"github.com/funtimecoding/soil/pkg/system"
	"path/filepath"
)

func cachePath() string {
	return filepath.Join(
		system.StorageDirectory("soil", false),
		"anthropic",
		"tokenizer.json",
	)
}
