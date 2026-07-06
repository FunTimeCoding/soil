package socket

import (
	"github.com/funtimecoding/soil/pkg/system"
	"path/filepath"
)

func directory() string {
	return filepath.Join(
		system.Home(),
		".local",
		"share",
		"goprocessd",
	)
}
