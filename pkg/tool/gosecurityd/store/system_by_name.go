package store

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/relational"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/store/system"
)

func (s *Store) SystemByName(name string) (*system.System, error) {
	var result system.System
	found, e := relational.FindOne(
		s.database.Where(map[string]any{"name": name}),
		&result,
	)

	if e != nil {
		return nil, fmt.Errorf("find system: %w", e)
	}

	if !found {
		return nil, not_found.New(constant.SystemKind, name)
	}

	return &result, nil
}
