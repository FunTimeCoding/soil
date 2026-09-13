package client

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/directory"
)

func addPerson(
	c *directory.Client,
	account string,
	name string,
) {
	add(
		c,
		fmt.Sprintf("uid=%s,ou=people,%s", account, c.Base()),
		map[string][]string{
			"objectClass":  {"inetOrgPerson"},
			"uid":          {account},
			"cn":           {name},
			"sn":           {"Person"},
			"mail":         {fmt.Sprintf("%s@example.test", account)},
			"userPassword": {fmt.Sprintf("%spassword", account)},
		},
	)
}
