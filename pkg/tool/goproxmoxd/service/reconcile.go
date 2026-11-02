package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_configured"
	proxmoxConstant "github.com/funtimecoding/soil/pkg/proxmox/constant"
	"github.com/funtimecoding/soil/pkg/proxmox/face"
	gonetboxdConstant "github.com/funtimecoding/soil/pkg/tool/gonetboxd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/types/intent"
	"github.com/luthermonson/go-proxmox"
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
			gonetboxdConstant.PowerIntentEnvironment,
		)
	}

	resources, e := guarded(
		s,
		hypervisor,
		func(c face.ProxmoxClient) (proxmox.ClusterResources, error) {
			return c.ClusterResources()
		},
	)

	if e != nil {
		return nil, e
	}

	machines, e := s.netboxMachines(x)

	if e != nil {
		return nil, e
	}

	result := intent.NewPlan(hypervisor)

	for _, r := range resources {
		if r.Type != proxmoxConstant.MachineType && r.Type != proxmoxConstant.ContainerType ||
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
