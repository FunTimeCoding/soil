package service

import (
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/armor"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/store/record"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/types/distinguished_name"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/types/issue_request"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/types/material"
)

func (s *Service) SignRequest(
	b *server.SigningRequestBody,
) (*record.Record, error) {
	a, e := s.authorityOf(b.Authority)

	if e != nil {
		return nil, e
	}

	q := armor.DecodeSigningRequest([]byte(b.Request))
	n := distinguished_name.From(q.Subject)
	r := issue_request.New()
	r.Kind = constant.CertificateKind(b.Kind)
	r.Name = n
	r.Host = host(q)
	r.ValidDay = validDay(b.ValidDay)

	return s.save(
		material.New(a.Sign(r, q.PublicKey), nil),
		r.Kind,
		"",
		b.Authority,
	)
}
