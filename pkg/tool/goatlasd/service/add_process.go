package service

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/collector/process"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/option"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/constant"
)

func (s *Service) addProcess(
	o *option.Atlas,
	l *logger.Logger,
	r face.Reporter,
) {
	c := o.Process()
	g := peer.New(constant.Identity.Name(), c.Address(), c, l, r)
	s.peers = append(s.peers, c)
	s.guards = append(s.guards, g)
	s.collectors = append(s.collectors, process.New(c, g, o.ProcessPlace))
}
