package loki

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/console/status"
	"github.com/funtimecoding/soil/pkg/console/status/option"
	"github.com/funtimecoding/soil/pkg/prometheus/loki/message"
	"github.com/funtimecoding/soil/pkg/time/constant"
)

func printLog(
	v []*message.Message,
	f *option.Format,
) {
	for _, m := range v {
		s := status.New(f).String(m.Time.Format(constant.DateMinute)).String(
			formatContent(m, f),
		).String(
			m.Stream,
		)
		console.Line(s.Format())
	}
}
