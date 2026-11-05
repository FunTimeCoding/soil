package goband

import (
	"github.com/funtimecoding/soil/pkg/tool/goband/plan"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/generated/client"
)

func describePower(
	v *client.Status,
	operation client.PowerRequestOperation,
	force bool,
	description string,
) *plan.Power {
	result := plan.NewPower(v.Instance, v.Power)

	if o := verdict(v, operation, force); o != nil && o.Refusal != nil {
		result.Refusal = o.Refusal

		return result
	}

	result.Description = description

	return result
}
