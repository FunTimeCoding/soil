package record

import (
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/armor"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/types/material"
)

func (r *Record) Material() *material.Material {
	c := armor.DecodeCertificate([]byte(r.Certificate))

	if r.Key == "" {
		return material.New(c, nil)
	}

	return material.New(c, armor.DecodeKey([]byte(r.Key)))
}
