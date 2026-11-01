package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/constant"
	"os"
	"path/filepath"
	"slices"
	"strings"
)

func (s *Service) CatchUpSearch(x context.Context) []error {
	entries, e := os.ReadDir(s.harbor)

	if e != nil {
		return []error{e}
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
		if x.Err() != nil {
			return nil
		}

		identifier := strings.TrimSuffix(
			i.Name(),
			constant.NotationLogExtension,
		)
		present[identifier] = true
		s.search.Append(identifier, filepath.Join(s.harbor, i.Name()))
		s.searchIndexed.Add(1)
	}

	indexed, f := s.search.Sessions()

	if f != nil {
		return []error{f}
	}

	for _, identifier := range indexed {
		if !present[identifier] {
			if _, g := s.search.DeleteSession(identifier); g != nil {
				return []error{g}
			}
		}
	}

	return nil
}
