package keystore

import (
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/armor"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/types/material"
	"path/filepath"
)

func Write(
	directory string,
	m *material.Material,
) {
	system.MakeDirectory(directory)
	system.WriteFile(
		filepath.Join(directory, constant.CertificateFile),
		armor.MarshalCertificate(m.Certificate),
		constant.CertificateMode,
	)
	system.WriteFile(
		filepath.Join(directory, constant.KeyFile),
		armor.MarshalKey(m.Key),
		constant.KeyMode,
	)
}
