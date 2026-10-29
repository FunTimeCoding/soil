package socket

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/system"
	"net"
	"os"
	"path/filepath"
)

func Listen(path string) net.Listener {
	system.MakeDirectory(filepath.Dir(path))

	if e := os.Remove(path); e != nil && !os.IsNotExist(e) {
		errors.PanicOnError(e)
	}

	result, f := net.Listen("unix", path)
	errors.PanicOnError(f)
	errors.PanicOnError(os.Chmod(path, 0o600))

	return result
}
