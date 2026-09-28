package outpost

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/face"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/gazetteer"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/model/placement"
	"github.com/funtimecoding/soil/pkg/tool/gooutpostd/generated/client"
	"time"
)

func collectOne(
	q context.Context,
	s *gazetteer.Gazetteer,
	o face.OutpostSource,
) ([]*placement.Placement, error) {
	host, e := o.GetHostWithResponse(q)

	if e != nil {
		return nil, e
	}

	if host.JSON200 == nil {
		return nil, unexpected.Format(
			constant.OutpostHostStatusFormat,
			host.HTTPResponse.StatusCode,
		)
	}

	where, okay := resolveHost(s, host.JSON200)

	if !okay {
		return nil, unexpected.Format(
			constant.PlaceUnknownFormat,
			host.JSON200.Hostname,
		)
	}

	services, f := o.ListServicesWithResponse(q, &client.ListServicesParams{})

	if f != nil {
		return nil, f
	}

	if services.JSON200 == nil {
		return nil, unexpected.Format(
			constant.OutpostServiceStatusFormat,
			services.HTTPResponse.StatusCode,
		)
	}

	var result []*placement.Placement
	now := time.Now()

	for _, v := range *services.JSON200 {
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
