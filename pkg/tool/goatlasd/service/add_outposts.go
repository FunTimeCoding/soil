package service

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/collector/outpost"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/option"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/types/target"
	"github.com/funtimecoding/soil/pkg/tool/gooutpostd/constant"
)

func (s *Service) addOutposts(
	o *option.Atlas,
	l *logger.Logger,
	r face.Reporter,
) {
	var targets []*target.Target

	for _, host := range o.OutpostHosts {
		c := o.Outpost(host)
		g := peer.New(
			join.Colon(constant.Identity.Name(), host),
			c.Address(),
			c,
			l,
			r,
		)
		s.peers = append(s.peers, c)
		s.guards = append(s.guards, g)
		targets = append(targets, target.New(host, c, g))
	}

	s.collectors = append(s.collectors, outpost.New(targets, l))
}
