package server_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/generated/server"
)

func RootBody() *server.AuthorityBody {
	country := constant.FixtureCountry
	province := constant.FixtureProvince
	organization := constant.FixtureOrganization

	return &server.AuthorityBody{
		Name:         constant.RootAuthority,
		Kind:         server.AuthorityKind(constant.KindRoot),
		CommonName:   constant.FixtureRootCommonName,
		Country:      &country,
		Province:     &province,
		Organization: &organization,
	}
}
