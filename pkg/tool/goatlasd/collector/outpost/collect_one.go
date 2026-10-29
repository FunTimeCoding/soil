package outpost

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/gazetteer"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/model/placement"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/types/target"
	"github.com/funtimecoding/soil/pkg/tool/gooutpostd/generated/client"
	"time"
)

func collectOne(
	q context.Context,
	s *gazetteer.Gazetteer,
	t *target.Target,
) ([]*placement.Placement, error) {
	var host *client.Host
	var services []client.Service
	e := t.Guard.Attempt(
		q,
		func(x context.Context) error {
			var f error

			if host, f = t.Client.Host(x); f != nil {
				return f
			}

			services, f = t.Client.Services(x, "")

			return f
		},
	)

	if e != nil {
		return nil, e
	}

	where, okay := resolveHost(s, host)

	if !okay {
		return nil, unexpected.Format(
			constant.PlaceUnknownFormat,
			host.Hostname,
		)
	}

	var result []*placement.Placement
	now := time.Now()

	for _, v := range services {
		if !deployed(v) {
			continue
		}

		result = append(
			result,
			placement.NewPackage(
				constant.SourceOutpost,
				constant.KindService,
				"",
				v.Name,
				text(v.Package),
				text(v.Version),
				where,
				now,
			),
		)
	}

	return result, nil
}
