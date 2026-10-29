package service

import (
	"errors"
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"io/fs"
	"os"
	"path/filepath"
)

func linkFiles(
	files []string,
	workdir string,
) error {
	for _, file := range files {
		if _, e := os.Stat(file); errors.Is(e, fs.ErrNotExist) {
			return not_found.New("file", file)
		} else if e != nil {
			return fmt.Errorf("stat file: %w", e)
		}

		if e := os.Symlink(
			file,
			filepath.Join(workdir, filepath.Base(file)),
		); e != nil {
			return fmt.Errorf("symlink file: %w", e)
		}
	}

	return nil
}
