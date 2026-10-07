package watcher

import (
	"crypto/sha256"
	"fmt"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/types/scanned_file"
	"os"
	"path/filepath"
	"strings"
)

func (w *Watcher) walkDirectory() []scanned_file.File {
	times, dirty := w.gitTimes()
	var result []scanned_file.File
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

				modifiedAt := i.ModTime()

				if t, okay := times[relative]; okay && !dirty[relative] {
					modifiedAt = t
				}

				result = append(
					result,
					scanned_file.File{
						Name: strings.TrimSuffix(
							filepath.Base(path),
							constant.MarkdownExtension,
						),
						Path:        relative,
						ContentHash: fmt.Sprintf("%x", sha256.Sum256(b)),
						Content:     string(b),
						ModifiedAt:  modifiedAt,
					},
				)

				return nil
			},
		),
	)

	return result
}
