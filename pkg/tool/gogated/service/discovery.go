package service

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gogated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogated/service/response"
)

func (s *Service) Discovery() *response.Discovery {
	d := response.NewDiscovery()
	d.Issuer = s.issuer
	d.AuthorizationEndpoint = fmt.Sprintf("%s/authorize", s.issuer)
	d.TokenEndpoint = fmt.Sprintf("%s/token", s.issuer)
	d.SigningKeysLocator = fmt.Sprintf("%s/jwks", s.issuer)
	d.RegistrationEndpoint = fmt.Sprintf("%s/register", s.issuer)
	d.EndSessionEndpoint = fmt.Sprintf("%s%s", s.issuer, constant.LogoutPath)
	d.ScopesSupported = []string{constant.ScopeOpenIdentity}
	d.ResponseTypesSupported = []string{constant.ResponseCode}
	d.GrantTypesSupported = []string{
		constant.GrantAuthorizationCode,
		constant.GrantRefreshToken,
	}
	d.SubjectTypesSupported = []string{constant.SubjectTypePublic}
	d.IdentifierTokenSigningAlgValuesSupported = []string{
		constant.AlgorithmRS256,
	}
	d.TokenEndpointAuthMethodsSupported = []string{
		constant.AuthMethodClientSecretBasic,
		constant.AuthMethodClientSecretPost,
		constant.AuthMethodNone,
	}
	d.CodeChallengeMethodsSupported = []string{constant.ChallengeMethodS256}

	return d
}
