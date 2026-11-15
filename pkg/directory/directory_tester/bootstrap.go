package directory_tester

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/directory"
)

func Bootstrap(c *directory.Client) {
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
}
