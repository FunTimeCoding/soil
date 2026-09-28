package credential

import "github.com/funtimecoding/soil/pkg/tool/gogated/constant"

func New(
	identifier string,
	secret string,
) *Credential {
	return &Credential{
		Identifier: identifier,
		Secret:     secret,
		Notice:     constant.SecretNotice,
	}
}
