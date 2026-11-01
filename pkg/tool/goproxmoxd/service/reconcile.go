package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_configured"
	"github.com/funtimecoding/soil/pkg/proxmox/constant"
	netbox "github.com/funtimecoding/soil/pkg/tool/gonetboxd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/types/intent"
	"slices"
	"strings"
)

func (s *Service) Reconcile(
	x context.Context,
	hypervisor string,
	full bool,
	write bool,
) (*intent.Plan, error) {
	if s.netbox == nil {
		return nil, not_configured.Format(
			"%s is not set",
			netbox.PowerIntentEnvironment,
		)
	}

	c, e := s.Client(hypervisor)

	if e != nil {
		return nil, e
	}

	resources, e := c.ClusterResources()

	if e != nil {
		return nil, e
	}

	machines, e := s.netboxMachines(x)

	if e != nil {
		return nil, e
	}

	result := intent.NewPlan(hypervisor)

	for _, r := range resources {
		if r.Type != constant.MachineType && r.Type != constant.ContainerType ||
			r.Template != 0 {
			continue
		}

		candidates := machines[r.Name]

		if len(candidates) == 0 {
			result.Unknown = append(result.Unknown, r.Name)

			continue
		}

		if len(candidates) > 1 {
			result.Ambiguous = append(result.Ambiguous, r.Name)

			continue
		}

		v := candidates[0]
		current := ""

		if v.Status != nil {
			current = *v.Status
		}

		if to, change := desiredStatus(r.Status, current, full); change {
			result.Changes = append(
				result.Changes,
				intent.NewChange(r.Name, r.Node, current, to),
			)
		}
	}

	slices.SortFunc(
		result.Changes,
		func(a, b *intent.Change) int { return strings.Compare(a.Name, b.Name) },
	)
	slices.Sort(result.Unknown)
	slices.Sort(result.Ambiguous)

	if !write {
		return result, nil
	}

	for _, change := range result.Changes {
		if f := s.writeGuest(x, change.Name, change.To); f != nil {
			return result, f
		}
	}

	result.Written = true

	return result, nil
}
