package user

import "github.com/funtimecoding/soil/pkg/tool/gogated/constant"

func NewDirectory(
	identifier string,
	unique string,
	account string,
	mail string,
	name string,
) *User {
	return &User{
		Identifier:          identifier,
		Mail:                mail,
		Source:              constant.SourceDirectory,
		DirectoryIdentifier: &unique,
		Account:             account,
		Name:                name,
	}
}
