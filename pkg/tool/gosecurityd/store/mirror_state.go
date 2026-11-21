package store

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/relational"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/store/mirror"
)

func (s *Store) MirrorState() (*mirror.Mirror, error) {
	result := mirror.Stub()

	if _, e := relational.FindOne(s.database.Order("id"), result); e != nil {
		return nil, fmt.Errorf("find score mirror: %w", e)
	}

	return result, nil
}
