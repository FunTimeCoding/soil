package service

import (
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/authority"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/store/record"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/types/distinguished_name"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/types/issue_request"
)

func (s *Service) CreateAuthority(
	b *server.AuthorityBody,
) (*record.Record, error) {
	live, e := s.store.Authority(b.Name)

	if e != nil {
		return nil, e
	}

	if live != nil {
		return nil, constant.ErrorConflict
	}

	kind := constant.CertificateKind(b.Kind)
	n := distinguished_name.New()
	n.CommonName = b.CommonName
	n.Country = text(b.Country)
	n.Province = text(b.Province)
	n.Organization = text(b.Organization)
	year := validYear(b.ValidYear, kind)

	if kind == constant.KindRoot {
		return s.save(authority.NewRoot(n, year).Material(), kind, b.Name, "")
	}

	parent, f := s.authorityOf(constant.RootAuthority)

	if f != nil {
		return nil, f
	}

	c, g := constraint(slice(b.PermittedDomain), slice(b.PermittedAddress))

	if g != nil {
		return nil, g
	}

	r := issue_request.New()
	r.Kind = kind
	r.Name = n
	r.Constraint = c
	r.ValidYear = year

	return s.save(parent.Issue(r), kind, b.Name, constant.RootAuthority)
}
