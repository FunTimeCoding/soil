package event

import (
	"github.com/funtimecoding/soil/pkg/kubernetes/check/event/option"
	kubernetes "github.com/funtimecoding/soil/pkg/kubernetes/constant"
	"github.com/funtimecoding/soil/pkg/kubernetes/types/native/event"
	"github.com/funtimecoding/soil/pkg/monitor/constant"
	item "github.com/funtimecoding/soil/pkg/monitor/item/constant"
	"github.com/funtimecoding/soil/pkg/monitor/report"
	"slices"
)

func printNotation(
	v []*event.Event,
	o *option.Event,
) {
	r := report.New()
	var relevant []*event.Event

	for _, e := range v {
		if !o.All && slices.Contains(
			kubernetes.IrrelevantEventReason,
			e.Reason,
		) {
			continue
		}

		relevant = append(relevant, e)
	}

	for _, e := range report.Trim(
		relevant,
		r,
		o.All,
		item.GoKevt,
	) {
		r.AddItem(
			item.GoKevt,
			e.MonitorIdentifier,
			constant.Warning,
			e.Reason,
			"",
			e.Create,
		)
	}

	r.Print()
}
