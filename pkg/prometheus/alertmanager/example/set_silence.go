package example

import (
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/constant"
	"github.com/funtimecoding/soil/pkg/tool/common"
)

func SetSilence() {
	common.Alertmanager().MustSimpleSilence(constant.NodeNotReady)
}
