package service

import (
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/errors"
	"os"
	"path/filepath"
	"slices"
	"strings"
)

func (s *Service) CatchUpSearch() {
	entries, e := os.ReadDir(s.harbor)

	if e != nil {
		s.reporter.CaptureException(e)

		return
	}

	var files []os.FileInfo

	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(
			entry.Name(),
			constant.NotationLogExtension,
		) {
			continue
		}

		if i, f := entry.Info(); f == nil {
			files = append(files, i)
		}
	}

	slices.SortFunc(
		files,
		func(a os.FileInfo, b os.FileInfo) int {
			return b.ModTime().Compare(a.ModTime())
		},
	)
	present := map[string]bool{}
	s.searchIndexed.Store(0)
	s.searchTotal.Store(int64(len(files)))

	for _, i := range files {
		identifier := strings.TrimSuffix(
			i.Name(),
			constant.NotationLogExtension,
		)
		present[identifier] = true
		s.search.Append(identifier, filepath.Join(s.harbor, i.Name()))
		s.searchIndexed.Add(1)
	}

	indexed, e := s.search.Sessions()
	errors.PanicOnError(e)

	for _, identifier := range indexed {
		if !present[identifier] {
			_, f := s.search.DeleteSession(identifier)
			errors.PanicOnError(f)
		}
	}
}
