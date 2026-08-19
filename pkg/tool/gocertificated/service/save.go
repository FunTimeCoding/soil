package service

import (
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/store/record"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/types/material"
)

func (s *Service) save(
	m *material.Material,
	kind constant.CertificateKind,
	name string,
	issuer string,
) (*record.Record, error) {
	r := record.New(kind, name, m)
	r.Issuer = issuer

	return r, s.store.Create(*r)
}
