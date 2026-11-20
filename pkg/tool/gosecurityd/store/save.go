package store

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/store/flag"
)

func (s *Store) Save(v *flag.Flag) error {
	if e := s.database.Save(v).Error; e != nil {
		return fmt.Errorf("save flag: %w", e)
	}

	return nil
}
