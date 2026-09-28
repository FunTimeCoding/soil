package service

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gogated/constant"
)

func (s *Service) Discovery() *DiscoveryDocument {
	return &DiscoveryDocument{
		Issuer:                s.issuer,
		AuthorizationEndpoint: fmt.Sprintf("%s/authorize", s.issuer),
		TokenEndpoint:         fmt.Sprintf("%s/token", s.issuer),
		SigningKeysLocator:    fmt.Sprintf("%s/jwks", s.issuer),
		RegistrationEndpoint:  fmt.Sprintf("%s/register", s.issuer),
		EndSessionEndpoint: fmt.Sprintf(
			"%s%s",
			s.issuer,
			constant.LogoutPath,
		),
		ScopesSupported:        []string{constant.ScopeOpenIdentity},
		ResponseTypesSupported: []string{constant.ResponseCode},
		GrantTypesSupported: []string{
			constant.GrantAuthorizationCode,
			constant.GrantRefreshToken,
		},
		SubjectTypesSupported: []string{constant.SubjectTypePublic},
		IdentifierTokenSigningAlgValuesSupported: []string{
			constant.AlgorithmRS256,
		},
		TokenEndpointAuthMethodsSupported: []string{
			constant.AuthMethodClientSecretBasic,
			constant.AuthMethodClientSecretPost,
			constant.AuthMethodNone,
		},
		CodeChallengeMethodsSupported: []string{constant.ChallengeMethodS256},
	}
}
