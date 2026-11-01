package service

import (
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
)

func (s *Service) queryNotice(e error) string {
	if e == nil {
		return ""
	}

	if unavailable.Is(e) {
		return constant.QueryAwayNotice
	}

	s.report(e)

	return constant.QueryFailedNotice
}
