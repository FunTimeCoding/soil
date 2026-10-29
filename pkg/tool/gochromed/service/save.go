package service

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/chromium/tab"
	"github.com/funtimecoding/soil/pkg/tool/gochromed/types/saved"
	"os"
	"path/filepath"
)

func (s *Service) save(
	t *tab.Tab,
	extension string,
	content []byte,
) (*saved.Saved, error) {
	path := filepath.Join(
		s.downloadDirectory,
		fmt.Sprintf("chrome_%s.%s", t.Identifier, extension),
	)

	if e := os.WriteFile(path, content, 0644); e != nil {
		return nil, fmt.Errorf("write file: %w", e)
	}

	return saved.New(path, t.Title, len(content)), nil
}
