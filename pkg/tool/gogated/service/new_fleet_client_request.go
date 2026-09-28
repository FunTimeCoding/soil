package service

import "github.com/funtimecoding/soil/pkg/tool/gogated/constant"

func NewFleetClientRequest(
	redirectLocators []string,
	scopes []string,
) *RegisterClientRequest {
	if len(scopes) == 0 {
		scopes = []string{constant.ScopeOpenIdentity, constant.ScopeOffline}
	}

	return NewRegisterClientRequest(
		redirectLocators,
		[]string{constant.GrantAuthorizationCode, constant.GrantRefreshToken},
		[]string{constant.ResponseCode},
		scopes,
		constant.AuthMethodClientSecretPost,
	)
}
