package daemon

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/option"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/peered"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/types/instance"
)

func peeredBuilder(
	o *option.Proxmox,
	l *logger.Logger,
	r face.Reporter,
) func(*instance.Instance) (*peered.Peered, error) {
	return func(i *instance.Instance) (*peered.Peered, error) {
		c, e := o.Hypervisor(i)

		if e != nil {
			return nil, e
		}

		return peered.New(
			c,
			peer.New(join.Colon(constant.InstancePeer, i.Name), i.Host, c, l, r),
		), nil
	}
}
