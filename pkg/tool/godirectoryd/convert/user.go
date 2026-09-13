package convert

import (
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/types/user"
)

func User(u *user.User) *server.User {
	return &server.User{
		Account:           u.Account,
		Name:              u.Name,
		Surname:           u.Surname,
		Mail:              u.Mail,
		Unique:            u.Unique,
		DistinguishedName: u.DistinguishedName,
	}
}
