package service

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/inventory/instance"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/option"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/service/peered"
)

func peeredBuilder(
	o *option.Postgres,
	l *logger.Logger,
	r face.Reporter,
) func(*instance.Instance) (*peered.Peered, error) {
	return func(i *instance.Instance) (*peered.Peered, error) {
		c, e := o.Client(i)

		if e != nil {
			return nil, e
		}

		return peered.New(
			c,
			peer.New(join.Colon(constant.InstancePeer, i.Name), i.Host, c, l, r),
		), nil
	}
}
