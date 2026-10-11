package technitium

import (
	"github.com/funtimecoding/soil/pkg/technitium/policy"
	"github.com/funtimecoding/soil/pkg/technitium/response"
	"slices"
)

func AllowUpdate(
	o *response.ZoneOptions,
	network string,
	p *policy.Policy,
) {
	o.Update = WidenAccess(o.Update)

	if !slices.Contains(o.UpdateNetwork, network) {
		o.UpdateNetwork = append(o.UpdateNetwork, network)
	}

	o.UpdatePolicies = policy.Upsert(o.UpdatePolicies, p)
}
