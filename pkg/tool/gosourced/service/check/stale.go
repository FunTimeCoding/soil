package check

import (
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"io/fs"
	"path/filepath"
	"strings"
)

func (s *Service) stale() string {
	if s.source == "" {
		return ""
	}

	var result string

	for _, d := range constant.LintSources {
		errors.PanicOnError(
			filepath.WalkDir(
				filepath.Join(s.source, d),
				func(
					p string,
					e fs.DirEntry,
					f error,
				) error {
					if f != nil || e.IsDir() ||
						!strings.HasSuffix(p, soil.GoExtension) ||
						strings.HasSuffix(p, soil.TestSuffix) {
						return nil
					}

					i, g := e.Info()

					if g == nil && i.ModTime().After(s.started) {
						relative, h := filepath.Rel(s.source, p)
						errors.PanicOnError(h)
						result = relative

						return fs.SkipAll
					}

					return nil
				},
			),
		)

		if result != "" {
			return result
		}
	}

	return ""
}
