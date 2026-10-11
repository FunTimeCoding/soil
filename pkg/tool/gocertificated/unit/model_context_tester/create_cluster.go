package model_context_tester

import (
	generative "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/constant"
)

func (o *Tester) CreateCluster() string {
	return o.Client.MustCallTool(
		constant.CreateAuthority,
		map[string]any{
			generative.ParameterName:     constant.FixtureClusterAuthority,
			constant.KindParameter:       string(constant.KindIntermediate),
			constant.CommonNameParameter: constant.FixtureIssuingCommonName,
			constant.DomainParameter: []any{
				constant.FixtureDomain,
				constant.FixtureInternalDomain,
				constant.FixtureLocalDomain,
			},
			constant.AddressParameter: []any{constant.FixtureAddress},
		},
	)
}
