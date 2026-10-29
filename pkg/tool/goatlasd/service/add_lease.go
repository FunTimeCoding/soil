package service

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/collector/lease"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/option"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/constant"
)

func (s *Service) addLease(
	o *option.Atlas,
	l *logger.Logger,
	r face.Reporter,
) {
	c := o.Lease()
	g := peer.New(constant.Identity.Name(), c.Address(), c, l, r)
	s.peers = append(s.peers, c)
	s.guards = append(s.guards, g)
	s.lease = lease.New(c, g)
}
