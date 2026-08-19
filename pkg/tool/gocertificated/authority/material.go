package authority

import "github.com/funtimecoding/soil/pkg/tool/gocertificated/types/material"

func (a *Authority) Material() *material.Material {
	return a.material
}
