package daemon

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	provision "github.com/funtimecoding/soil/pkg/provision/face"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/service/feed"
	"github.com/funtimecoding/soil/pkg/worker"
)

func (d *Daemon) withFeed(
	a provision.Client,
	l *logger.Logger,
	r face.Reporter,
) {
	d.ansible = a
	d.ansibleGuard = peer.New(constant.AnsiblePeer, a.Address(), a, l, r)
	d.feed = feed.New(a, d.ansibleGuard, d.system, l)
	d.feedClock = worker.New(d.feed.Run, l, r, constant.ProbeInterval)
}
