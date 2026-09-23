//go:build ci

package client

import (
	"fmt"
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
