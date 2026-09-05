package publish_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/generated/server"
	"testing"
)

func (o *Tester) CreateRoot(t *testing.T) {
	t.Helper()
	country := constant.FixtureCountry
	province := constant.FixtureProvince
	organization := constant.FixtureOrganization
	_, e := o.Server.Service.CreateAuthority(
		&server.AuthorityBody{
			Name:         constant.RootAuthority,
			Kind:         server.AuthorityKind(constant.KindRoot),
			CommonName:   constant.FixtureRootCommonName,
			Country:      &country,
			Province:     &province,
			Organization: &organization,
		},
	)
	assert.FatalOnError(t, e)
}
