package score_scale

import "github.com/funtimecoding/soil/pkg/math/range_mapping"

func (c *Scale) Mapping() []*range_mapping.Mapping {
	return c.mapping
}
