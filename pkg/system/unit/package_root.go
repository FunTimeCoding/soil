package unit

import (
	"github.com/funtimecoding/soil/pkg/system"
	"path/filepath"
)

func packageRoot() string {
	return filepath.Join(system.WorkDirectory(), "goexample_1.0.0-1_amd64")
}
