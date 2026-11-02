package notify

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/prometheus"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/tool/common"
)

func Notify() {
	c := common.Alertmanager()
	defer c.Close()
	p := prometheus.NewEnvironment()
	defer p.Close()
	s := &State{}
	stop := make(chan struct{})
	go worker(stop, c, p, s)
	system.KillSignalBlock()
	close(stop)
	console.Line("Stop")
}
