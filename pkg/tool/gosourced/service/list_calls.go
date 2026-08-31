package service

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/source/resolve"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/result"
	"go/types"
	"sort"
	"strings"
)

func (s *Service) ListCalls(
	directory string,
	region string,
	limit int,
) (*output.Results, *result.Inventory, error) {
	r := output.NewResultsWithDirectory(directory)
	all, _, e := resolve.LoadPackages(directory, "./...")

	if e != nil {
		return nil, nil, e
	}

	counts := map[string]int{}
	found := false

	for _, p := range all {
		if !strings.HasPrefix(p.PkgPath, region) {
			continue
		}

		found = true

		for _, use := range p.TypesInfo.Uses {
			function, okay := use.(*types.Func)

			if !okay || function.Pkg() == nil {
				continue
			}

			if strings.HasPrefix(function.Pkg().Path(), region) {
				continue
			}

			counts[function.FullName()]++
		}
	}

	if !found {
		r.AddConcern(
			concern.NewFile(
				"validation",
				fmt.Sprintf("no packages under region: %s", region),
				"",
				false,
			),
		)

		return r, nil, nil
	}

	var calls []*result.Call

	for name, count := range counts {
		calls = append(calls, result.NewCall(name, count))
	}

	sort.Slice(
		calls,
		func(i, j int) bool {
			if calls[i].Count != calls[j].Count {
				return calls[i].Count > calls[j].Count
			}

			return calls[i].Name < calls[j].Name
		},
	)
	total := len(calls)
	more := 0

	if limit > 0 && total > limit {
		more = total - limit
		calls = calls[:limit]
	}

	return r, result.NewInventory(region, total, more, calls), nil
}
