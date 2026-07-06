package parse

import (
	"github.com/funtimecoding/soil/pkg/prometheus/constant"
	"github.com/funtimecoding/soil/pkg/prometheus/result/generic"
	"github.com/prometheus/common/model"
	"log"
)

func Generic(v model.Value) []*generic.Result {
	var result []*generic.Result

	switch t := v.Type(); t {
	case model.ValMatrix:
		for _, e := range v.(model.Matrix) {
			for _, inner := range e.Values {
				result = append(
					result,
					generic.New(
						constant.Matrix,
						e.Metric.String(),
						inner.Timestamp.Time(),
						inner.Value.String(),
					),
				)
			}
		}
	case model.ValVector:
		for _, e := range v.(model.Vector) {
			result = append(
				result,
				generic.New(
					constant.Vector,
					e.Metric.String(),
					e.Timestamp.Time(),
					e.Value.String(),
				),
			)
		}
	case model.ValScalar:
		s := v.(*model.Scalar)
		result = append(
			result,
			generic.New(
				constant.Scalar,
				s.String(),
				s.Timestamp.Time(),
				s.Value.String(),
			),
		)
	case model.ValString:
		s := v.(*model.String)
		result = append(
			result,
			generic.New(
				constant.String,
				s.String(),
				s.Timestamp.Time(),
				s.Value,
			),
		)
	default:
		log.Panicf("unexpected type: %d", t)
	}

	return result
}
