package service

import (
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/armor"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/store/record"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/types/distinguished_name"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/types/issue_request"
)

func (s *Service) IssueCertificate(
	b *server.CertificateBody,
) (*record.Record, string, error) {
	a, e := s.authorityOf(b.Authority)

	if e != nil {
		return nil, "", e
	}

	n := distinguished_name.New()
	n.CommonName = b.CommonName
	r := issue_request.New()
	r.Kind = constant.CertificateKind(b.Kind)
	r.Name = n
	r.Host = slice(b.Host)
	r.ValidDay = validDay(b.ValidDay)
	m := a.Issue(r)
	result, f := s.save(m, r.Kind, "", b.Authority)

	if f != nil {
		return nil, "", f
	}

	return result, string(armor.MarshalKey(m.Key)), nil
}
