package register_client

import "github.com/funtimecoding/soil/pkg/tool/gogated/constant"

func NewFleetRequest(
	redirectLocators []string,
	scopes []string,
) *Request {
	if len(scopes) == 0 {
		scopes = []string{constant.ScopeOpenIdentity, constant.ScopeOffline}
	}

	return NewRequest(
		redirectLocators,
		[]string{constant.GrantAuthorizationCode, constant.GrantRefreshToken},
		[]string{constant.ResponseCode},
		scopes,
		constant.AuthMethodClientSecretPost,
	)
}
