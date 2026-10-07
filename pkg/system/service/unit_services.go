package service

import (
	"github.com/funtimecoding/soil/pkg/system/constant"
	"github.com/funtimecoding/soil/pkg/system/types/service"
)

func (c *Client) unitServices() []*service.Service {
	units := ParseUnits(
		output(
			constant.Systemctl,
			constant.SystemctlList,
			constant.SystemctlServiceType,
			constant.SystemctlAll,
			constant.SystemctlNoLegend,
			constant.SystemctlNoPager,
			constant.SystemctlPlain,
		),
	)
	path := unitPaths()
	owner := ParseSearch(
		output(
			constant.DpkgQuery,
			append(
				[]string{constant.DpkgSearch},
				searchPaths(path)...,
			)...,
		),
	)
	names := make(map[string]string)

	for unit, p := range path {
		names[unit] = OwnerOf(owner, p)
	}

	policies := ParsePolicy(
		output(
			constant.AptCache,
			append(
				[]string{constant.AptCachePolicy},
				values(names)...,
			)...,
		),
	)
	manual := ParseManual(output(constant.AptMark, constant.AptMarkManual))
	var result []*service.Service

	for unit, state := range units {
		result = append(
			result,
			service.NewUnit(
				unit,
				state,
				path[unit],
				names[unit],
				policies,
				manual,
			),
		)
	}

	return result
}
