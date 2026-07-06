package watcher

import (
	"crypto/sha256"
	"fmt"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/errors"
	"os"
	"path/filepath"
	"strings"
)

func (w *Watcher) walkDirectory() []scannedFile {
	var result []scannedFile
	errors.PanicOnError(
		filepath.Walk(
			w.seedDirectory,
			func(
				path string,
				i os.FileInfo,
				e error,
			) error {
				if e != nil {
					return nil
				}

				if i.IsDir() {
					return nil
				}

				if !strings.HasSuffix(path, constant.MarkdownExtension) {
					return nil
				}

				b, f := os.ReadFile(path)

				if f != nil {
					return nil
				}

				relative, g := filepath.Rel(w.seedDirectory, path)

				if g != nil {
					return nil
				}

				result = append(
					result,
					scannedFile{
						name: strings.TrimSuffix(
							filepath.Base(path),
							constant.MarkdownExtension,
						),
						path:        relative,
						contentHash: fmt.Sprintf("%x", sha256.Sum256(b)),
						content:     string(b),
					},
				)

				return nil
			},
		),
	)

	return result
}
