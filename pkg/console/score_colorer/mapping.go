package score_colorer

import "github.com/funtimecoding/soil/pkg/math/range_mapping"

func (c *Colorer) Mapping() []*range_mapping.Mapping {
	return c.mapping
}
