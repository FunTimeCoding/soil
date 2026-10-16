package host_mapping

import (
	provisionConstant "github.com/funtimecoding/soil/pkg/provision/constant"
	"github.com/funtimecoding/soil/pkg/tool/goterraformd/types/plan"
)

func (x *Index) ownHost(
	c *plan.ResourceChange,
	a string,
) string {
	values := c.Change.After

	if a == provisionConstant.HostDelete {
		values = c.Change.Before
	}

	s, _ := values[x.attribute[c.Type]].(string)

	return s
}
