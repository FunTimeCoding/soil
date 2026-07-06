package claude

import (
	"github.com/funtimecoding/soil/pkg/system"
	"path/filepath"
)

func New() *Client {
	return &Client{
		base: filepath.Join(
			system.Home(),
			".local",
			"share",
			"goclauded",
			"session",
		),
	}
}
