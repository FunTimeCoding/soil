package claude

import (
	"github.com/funtimecoding/soil/pkg/system"
	"path/filepath"
)

func New() *Client {
	return NewDirectory(
		filepath.Join(
			system.StorageDirectory("goclauded", false),
			"session",
		),
	)
}
