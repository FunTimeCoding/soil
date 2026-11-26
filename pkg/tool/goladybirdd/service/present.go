package service

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goladybirdd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goladybirdd/types/view"
	"os"
	"path/filepath"
)

func (s *Service) present(
	tab string,
	text string,
) (*view.View, error) {
	result := view.New(text)

	if len(text) <= constant.SnapshotLimit {
		return result, nil
	}

	result.Path = filepath.Join(
		s.downloadDirectory,
		fmt.Sprintf("ladybird_%s_snapshot.txt", tab),
	)

	if e := os.WriteFile(result.Path, []byte(text), 0644); e != nil {
		return nil, fmt.Errorf("write file: %w", e)
	}

	result.Text = ""

	return result, nil
}
