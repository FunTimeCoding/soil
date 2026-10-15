package observation_tester

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/web/constant"
)

func (o *Observation) Codes() map[string]uint64 {
	families, e := o.registry.Gather()
	errors.PanicOnError(e)
	result := map[string]uint64{}

	for _, f := range families {
		for _, m := range f.GetMetric() {
			l := map[string]string{}

			for _, v := range m.GetLabel() {
				l[v.GetName()] = v.GetValue()
			}

			result[join.Space(
				l[constant.MetricCodeLabel],
				l[constant.MetricMethodLabel],
			)] = m.GetHistogram().GetSampleCount()
		}
	}

	return result
}
