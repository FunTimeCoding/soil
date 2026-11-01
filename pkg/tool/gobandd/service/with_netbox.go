package service

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/constant"
	netbox "github.com/funtimecoding/soil/pkg/tool/gonetboxd/face"
)

func (s *Service) WithNetbox(
	n netbox.Client,
	l *logger.Logger,
	r face.Reporter,
) *Service {
	s.netbox = n
	s.netboxGuard = peer.New(constant.Identity.Name(), n.Address(), n, l, r)
	s.logger = l
	s.reporter = r

	return s
}
