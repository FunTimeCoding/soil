package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store/record"
)

func (s *Service) index(m *record.Memory) {
	s.report(s.syncIndex(context.Background(), m))
}
