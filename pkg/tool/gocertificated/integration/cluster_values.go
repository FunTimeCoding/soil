package integration

import (
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/constant"
	"net/url"
)

func clusterValues() url.Values {
	return url.Values{
		"name":                       {constant.FixtureClusterAuthority},
		constant.KindParameter:       {string(constant.KindIntermediate)},
		constant.CommonNameParameter: {constant.FixtureIssuingCommonName},
		constant.DomainParameter:     {constant.FixtureDomain},
		constant.AddressParameter:    {constant.FixtureAddress},
	}
}
