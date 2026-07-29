package example

import (
	"github.com/funtimecoding/soil/pkg/prometheus/constant"
	"github.com/funtimecoding/soil/pkg/tool/common"
)

func SetSilence() {
	common.Alertmanager().MustSimpleSilence(constant.NodeNotReady)
}
