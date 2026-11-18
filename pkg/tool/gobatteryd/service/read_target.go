package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/nut/detached"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/gobatteryd/types/reading"
	"github.com/funtimecoding/soil/pkg/tool/gobatteryd/types/target"
)

func readTarget(t *target.Target) ([]*reading.Reading, error) {
	return peer.Guarded(
		context.Background(),
		t.Guard,
		func(x context.Context) ([]*reading.Reading, error) {
			batteries, e := t.Client.Batteries(x)

			if e != nil {
				return nil, e
			}

			var result []*reading.Reading

			for _, b := range batteries {
				v, f := t.Client.Variables(x, b)

				if detached.Is(f) {
					result = append(result, reading.NewDetached(t.Server, b))

					continue
				}

				if f != nil {
					return nil, f
				}

				result = append(result, reading.New(t.Server, b, v))
			}

			return result, nil
		},
	)
}
