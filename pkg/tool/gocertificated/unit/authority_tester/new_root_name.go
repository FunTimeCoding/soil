package authority_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/types/distinguished_name"
)

func NewRootName() *distinguished_name.Name {
	n := distinguished_name.New()
	n.Country = constant.FixtureCountry
	n.Province = constant.FixtureProvince
	n.Organization = constant.FixtureOrganization
	n.CommonName = constant.FixtureRootCommonName

	return n
}
