package publish_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/generated/server"
	"testing"
)

func (o *Tester) CreateCluster(t *testing.T) {
	t.Helper()
	domain := []string{
		constant.FixtureDomain,
		constant.FixtureInternalDomain,
		constant.FixtureLocalDomain,
	}
	address := []string{constant.FixtureAddress}
	_, e := o.Server.Service.CreateAuthority(
		&server.AuthorityBody{
			Name:             constant.FixtureClusterAuthority,
			Kind:             server.AuthorityKind(constant.KindIntermediate),
			CommonName:       constant.FixtureIssuingCommonName,
			PermittedDomain:  &domain,
			PermittedAddress: &address,
		},
	)
	assert.FatalOnError(t, e)
}
