package service

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/gosentryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosentryd/option"
)

func New(
	o *option.Server,
	l *logger.Logger,
	r face.Reporter,
) *Service {
	c := o.Source()

	return &Service{
		client:       c,
		guard:        peer.New(constant.SentryPeer, o.Host, c, l, r),
		organization: o.Organization,
	}
}
