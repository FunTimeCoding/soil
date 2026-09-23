package server_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/generated/server"
)

func ClusterBody() *server.AuthorityBody {
	domain := []string{
		constant.FixtureDomain,
		constant.FixtureInternalDomain,
		constant.FixtureLocalDomain,
	}
	address := []string{constant.FixtureAddress}

	return &server.AuthorityBody{
		Name:             constant.FixtureClusterAuthority,
		Kind:             server.AuthorityKind(constant.KindIntermediate),
		CommonName:       constant.FixtureIssuingCommonName,
		PermittedDomain:  &domain,
		PermittedAddress: &address,
	}
}
