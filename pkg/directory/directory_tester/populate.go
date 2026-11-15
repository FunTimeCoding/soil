package directory_tester

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/directory"
)

func Populate(c *directory.Client) {
	Bootstrap(c)
	addPerson(c, "member", "Member Person")
	addPerson(c, "outsider", "Outsider Person")
	add(
		c,
		fmt.Sprintf("cn=gate,ou=groups,%s", c.Base()),
		map[string][]string{
			"objectClass": {"posixGroup"},
			"cn":          {"gate"},
			"gidNumber":   {"5000"},
			"memberUid":   {"member"},
		},
	)
}
