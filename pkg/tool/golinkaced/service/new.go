package service

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/golinkaced/constant"
	"github.com/funtimecoding/soil/pkg/tool/golinkaced/option"
)

func New(
	o *option.LinkAce,
	l *logger.Logger,
	r face.Reporter,
) *Service {
	c := o.Source()

	return &Service{
		client: c,
		guard:  peer.New(constant.LinkAcePeer, c.Address(), c, l, r),
	}
}
