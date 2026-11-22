package worker

import (
	"github.com/funtimecoding/soil/pkg/tool/gobatteryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gobatteryd/types/reading"
	"strconv"
	"strings"
)

func (w *Worker) record(r *reading.Reading) {
	driver := 0.0

	if r.DriverConnected {
		driver = 1
	}

	w.driver.WithLabelValues(r.Server, r.Battery).Set(driver)

	if !r.Attached {
		w.attached.WithLabelValues(r.Server, r.Battery).Set(0)

		return
	}

	w.attached.WithLabelValues(r.Server, r.Battery).Set(1)

	for variable, g := range w.measures {
		v, found := r.Variables[variable]

		if !found {
			continue
		}

		if f, e := strconv.ParseFloat(v, 64); e == nil {
			g.WithLabelValues(r.Server, r.Battery).Set(f)
		}
	}

	for _, flag := range strings.Fields(r.Variables[constant.StatusVariable]) {
		w.status.WithLabelValues(r.Server, r.Battery, statusWord(flag)).Set(1)
	}

	w.information.WithLabelValues(
		r.Server,
		r.Battery,
		strings.TrimSpace(r.Variables[constant.ModelVariable]),
		strings.TrimSpace(r.Variables[constant.ManufacturerVariable]),
		strings.TrimSpace(r.Variables[constant.SerialVariable]),
	).Set(1)
}
