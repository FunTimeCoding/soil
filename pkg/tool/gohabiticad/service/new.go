package service

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/constant"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/option"
)

func New(
	o *option.Habitica,
	l *logger.Logger,
	r face.Reporter,
) *Service {
	c := o.Source()

	return &Service{
		client: c,
		guard:  peer.New(constant.HabiticaPeer, o.Host, c, l, r),
	}
}
