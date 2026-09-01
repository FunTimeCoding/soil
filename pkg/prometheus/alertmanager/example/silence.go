package example

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/tool/common"
)

func Silence() {
	f := constant.ColorFormat.Copy().Tag(constant.TagState)

	for _, a := range common.Alertmanager().MustSilences(true) {
		console.Line(a.Format(f))
	}
}
