package integration

import (
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/constant"
	"net/url"
)

func rootValues() url.Values {
	return url.Values{
		"name":                         {constant.RootAuthority},
		constant.KindParameter:         {string(constant.KindRoot)},
		constant.CommonNameParameter:   {constant.FixtureRootCommonName},
		constant.CountryParameter:      {constant.FixtureCountry},
		constant.ProvinceParameter:     {constant.FixtureProvince},
		constant.OrganizationParameter: {constant.FixtureOrganization},
	}
}
