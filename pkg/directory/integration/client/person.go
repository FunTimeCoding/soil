//go:build ci

package client

import "github.com/funtimecoding/soil/pkg/directory/constant"

func person(account string) map[string][]string {
	return map[string][]string{
		"objectClass":             {"inetOrgPerson"},
		constant.AccountAttribute: {account},
		constant.NameAttribute:    {account},
		"sn":                      {account},
	}
}
