package dependency

import "slices"

func (d *Dependency) HasConcern(s string) bool {
	return slices.Contains(d.concern, s)
}
