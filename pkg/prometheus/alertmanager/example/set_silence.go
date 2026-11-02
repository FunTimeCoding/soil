package example

import (
	"github.com/funtimecoding/soil/pkg/prometheus"
	"github.com/funtimecoding/soil/pkg/prometheus/constant"
	"github.com/funtimecoding/soil/pkg/tool/common"
)

func SetSilence() {
	c := common.Alertmanager()
	defer c.Close()
	p := prometheus.NewEnvironment()
	defer p.Close()
	c.MustSimpleSilence(constant.NodeNotReady, p.MustRules())
}
