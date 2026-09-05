package model_context_tester

import (
	generative "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/constant"
)

func rootArgument() map[string]any {
	return map[string]any{
		generative.ParameterName:       constant.RootAuthority,
		constant.KindParameter:         string(constant.KindRoot),
		constant.CommonNameParameter:   constant.FixtureRootCommonName,
		constant.CountryParameter:      constant.FixtureCountry,
		constant.ProvinceParameter:     constant.FixtureProvince,
		constant.OrganizationParameter: constant.FixtureOrganization,
	}
}
