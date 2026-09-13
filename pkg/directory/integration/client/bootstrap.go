package client

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/directory"
)

func bootstrap(c *directory.Client) {
	add(
		c,
		c.Base(),
		map[string][]string{
			"objectClass": {"dcObject", "organization"},
			"dc":          {"example"},
			"o":           {"example"},
		},
	)
	add(
		c,
		fmt.Sprintf("ou=people,%s", c.Base()),
		map[string][]string{
			"objectClass": {"organizationalUnit"},
			"ou":          {"people"},
		},
	)
	add(
		c,
		fmt.Sprintf("ou=groups,%s", c.Base()),
		map[string][]string{
			"objectClass": {"organizationalUnit"},
			"ou":          {"groups"},
		},
	)
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
