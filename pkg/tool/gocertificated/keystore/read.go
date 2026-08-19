package keystore

import (
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/armor"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/types/material"
)

func Read(directory string) *material.Material {
	return material.New(
		armor.DecodeCertificate(
			system.ReadBytes(directory, constant.CertificateFile),
		),
		armor.DecodeKey(system.ReadBytes(directory, constant.KeyFile)),
	)
}
