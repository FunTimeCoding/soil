package convert

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/prometheus/query_result"
	"github.com/prometheus/common/model"
	"time"
)

func QueryResult(r *query_result.Result) *SlimQueryResult {
	result := &SlimQueryResult{
		Type:     r.Value.Type().String(),
		Warnings: r.Warnings,
	}

	switch t := r.Value.(type) {
	case model.Vector:
		for _, s := range t {
			result.Vector = append(
				result.Vector,
				SlimSample{
					Metric:    metricToMap(s.Metric),
					Value:     s.Value.String(),
					Timestamp: s.Timestamp.Time().Format(time.RFC3339),
				},
			)
		}
	case model.Matrix:
		for _, s := range t {
			stream := SlimStream{
				Metric: metricToMap(s.Metric),
			}

			for _, p := range s.Values {
				stream.Values = append(
					stream.Values,
					SlimPoint{
						Value:     p.Value.String(),
						Timestamp: p.Timestamp.Time().Format(time.RFC3339),
					},
				)
			}

			result.Matrix = append(result.Matrix, stream)
		}
	case *model.Scalar:
		result.Scalar = &SlimPoint{
			Value:     t.Value.String(),
			Timestamp: t.Timestamp.Time().Format(time.RFC3339),
		}
	default:
		panic(fmt.Sprintf("unexpected query result type: %s", r.Value.Type()))
	}

	return result
}
