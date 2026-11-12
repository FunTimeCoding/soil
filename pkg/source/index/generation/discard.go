package generation

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/source/constant"
	"os"
	"path/filepath"
	"strings"
)

func discard(
	root string,
	name string,
) {
	path := filepath.Join(root, name)

	if !strings.HasPrefix(name, constant.GenerationTrash) {
		trash := filepath.Join(
			root,
			fmt.Sprintf("%s%d-%s", constant.GenerationTrash, os.Getpid(), name),
		)

		if os.Rename(path, trash) != nil {
			return
		}

		path = trash
	}

	errors.LogOnError(os.RemoveAll(path))
}
