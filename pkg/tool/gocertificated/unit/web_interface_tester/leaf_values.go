package web_interface_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/constant"
	"net/url"
)

func LeafValues() url.Values {
	return url.Values{
		constant.AuthorityParameter:  {constant.FixtureClusterAuthority},
		constant.KindParameter:       {string(constant.KindServer)},
		constant.CommonNameParameter: {constant.FixtureCommonName},
		constant.HostParameter:       {constant.FixtureHost},
	}
}
