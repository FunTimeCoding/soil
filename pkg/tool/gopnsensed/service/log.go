package service

import (
	"github.com/funtimecoding/soil/pkg/opnsense/log_entry"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/constant"
)

func (s *Service) Log(limit int) ([]*log_entry.Entry, error) {
	if limit == 0 {
		limit = constant.DefaultLogLimit
	}

	return s.source.Log(limit)
}
