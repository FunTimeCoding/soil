package user

import "github.com/funtimecoding/soil/pkg/tool/gogated/constant"

func New(
	identifier string,
	mail string,
	passwordHash string,
) *User {
	return &User{
		Identifier:   identifier,
		Mail:         mail,
		PasswordHash: passwordHash,
		Source:       constant.SourceLocal,
	}
}
