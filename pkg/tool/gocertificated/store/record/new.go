package record

import (
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/armor"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/types/material"
)

func New(
	kind constant.CertificateKind,
	name string,
	m *material.Material,
) *Record {
	r := Stub()
	r.Serial = m.Certificate.SerialNumber.Text(constant.SerialBase)
	r.Kind = string(kind)
	r.Name = name
	r.CommonName = m.Certificate.Subject.CommonName
	r.Certificate = string(armor.MarshalCertificate(m.Certificate))
	r.Start = m.Certificate.NotBefore
	r.End = m.Certificate.NotAfter

	if constant.AuthorityKind[kind] {
		r.Key = string(armor.MarshalKey(m.Key))
	}

	return r
}
