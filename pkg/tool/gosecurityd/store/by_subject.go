package store

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/relational"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/store/flag"
)

func (s *Store) BySubject(subject string) (*flag.Flag, error) {
	var result flag.Flag
	found, e := relational.FindOne(
		s.database.Where(map[string]any{constant.FlagSubjectColumn: subject}),
		&result,
	)

	if e != nil {
		return nil, fmt.Errorf("find flag: %w", e)
	}

	if !found {
		return nil, not_found.New(constant.FlagKind, subject)
	}

	return &result, nil
}
