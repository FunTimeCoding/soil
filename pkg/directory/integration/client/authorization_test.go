//go:build ci

package client

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestServiceAccountWritesUsers(t *testing.T) {
	c := service(t)
	name := fmt.Sprintf("uid=authorized,ou=people,%s", c.Base())
	assert.FatalOnError(
		t,
		c.Add(
			name,
			map[string][]string{
				"objectClass": {"inetOrgPerson"},
				"uid":         {"authorized"},
				"cn":          {"Authorized Person"},
				"sn":          {"Person"},
			},
		),
	)

	defer func() { assert.FatalOnError(t, c.Delete(name)) }()
	assert.FatalOnError(t, c.SetPassword(name, "authorizedpassword"))
	entry, e := c.Authenticate("authorized", "authorizedpassword")
	assert.FatalOnError(t, e)
	assert.String(t, "authorized", entry.Account)
}

func TestServiceAccountRefusedOutsideScope(t *testing.T) {
	c := service(t)
	e := c.Add(
		fmt.Sprintf("ou=elsewhere,%s", c.Base()),
		map[string][]string{
			"objectClass": {"organizationalUnit"},
			"ou":          {"elsewhere"},
		},
	)

	if e == nil {
		assert.FatalOnError(
			t,
			c.Delete(fmt.Sprintf("ou=elsewhere,%s", c.Base())),
		)
		t.Fatalf("expected the service account to be refused outside its scope")
	}
}

func TestServiceAccountRefusedOnConfiguration(t *testing.T) {
	c := service(t)
	e := c.Modify(
		"olcDatabase={1}mdb,cn=config",
		map[string][]string{"olcDbIndex": {"uid eq"}},
	)

	if e == nil {
		t.Fatalf("expected the service account to be refused on cn=config")
	}
}
