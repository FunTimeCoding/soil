package authority

import "github.com/funtimecoding/soil/pkg/tool/gocertificated/types/material"

func New(m *material.Material) *Authority {
	return &Authority{material: m}
}
