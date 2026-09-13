//go:build ci

package client

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/directory"
	"github.com/funtimecoding/soil/pkg/directory/constant"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"testing"
)

func service(t *testing.T) *directory.Client {
	t.Helper()
	password := environment.Optional("LDAP_SERVICE_PASSWORD")

	if password == "" {
		t.Skip("LDAP_SERVICE_PASSWORD unset")
	}

	c := directory.NewEnvironment()

	return directory.New(
		environment.Required(constant.HostEnvironment),
		c.Base(),
	).
		WithPort(environment.RequiredInteger(constant.PortEnvironment)).
		WithInsecure().
		WithCredential(
			fmt.Sprintf("cn=directory,ou=services,%s", c.Base()),
			password,
		).
		WithFilter(
			environment.Required(constant.UserFilterEnvironment),
			environment.Required(constant.GroupFilterEnvironment),
		)
}

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
