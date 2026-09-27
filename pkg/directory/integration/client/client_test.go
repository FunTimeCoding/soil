//go:build ci

package client

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/directory"
	"github.com/funtimecoding/soil/pkg/directory/constant"
	"github.com/funtimecoding/soil/pkg/errors/conflict"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	bootstrap(directory.NewEnvironment())
	os.Exit(m.Run())
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

func TestAuthenticateMember(t *testing.T) {
	entry, e := directory.NewEnvironment().Authenticate(
		"member",
		"memberpassword",
	)
	assert.FatalOnError(t, e)
	assert.String(t, "member", entry.Account)
	assert.String(t, "member@example.test", entry.Mail)
	assert.String(t, "Member Person", entry.Name)
	assert.String(
		t,
		"uid=member,ou=people,dc=example,dc=test",
		entry.DistinguishedName,
	)

	if entry.Unique == "" {
		t.Errorf("expected a unique identifier")
	}
}

func TestAuthenticateByMail(t *testing.T) {
	entry, e := directory.NewEnvironment().Authenticate(
		"member@example.test",
		"memberpassword",
	)
	assert.FatalOnError(t, e)
	assert.String(t, "member", entry.Account)
}

func TestAuthenticateWrongPassword(t *testing.T) {
	_, e := directory.NewEnvironment().Authenticate("member", "wrongpassword")
	assert.True(t, validation.Is(e))
}

func TestAuthenticateUnknown(t *testing.T) {
	_, e := directory.NewEnvironment().Authenticate("ghost", "anypassword")
	assert.True(t, not_found.Is(e))
}

func TestAuthenticateFilterInjection(t *testing.T) {
	_, e := directory.NewEnvironment().Authenticate("*", "memberpassword")
	assert.True(t, not_found.Is(e))
}

func TestInGroupMember(t *testing.T) {
	member, e := directory.NewEnvironment().InGroup("member")
	assert.FatalOnError(t, e)
	assert.True(t, member)
}

func TestInGroupOutsider(t *testing.T) {
	member, e := directory.NewEnvironment().InGroup("outsider")
	assert.FatalOnError(t, e)
	assert.False(t, member)
}

func TestUniqueIdentifierStable(t *testing.T) {
	c := directory.NewEnvironment()
	first, e := c.Authenticate("member", "memberpassword")
	assert.FatalOnError(t, e)
	second, f := c.Authenticate("member", "memberpassword")
	assert.FatalOnError(t, f)
	assert.String(t, first.Unique, second.Unique)
}

func TestAddModifyDelete(t *testing.T) {
	c := directory.NewEnvironment()
	name := distinguished("written")
	assert.FatalOnError(t, c.Add(name, person("written")))
	found, e := c.Search(
		"(uid=written)",
		[]string{constant.NameAttribute, "sn"},
	)
	assert.FatalOnError(t, e)
	assert.Count(t, 1, found)
	assert.String(t, "written", found[0].Attributes[constant.NameAttribute][0])
	assert.FatalOnError(
		t,
		c.Modify(
			name,
			map[string][]string{constant.NameAttribute: {"Renamed Person"}},
		),
	)
	after, f := c.Search("(uid=written)", []string{constant.NameAttribute})
	assert.FatalOnError(t, f)
	assert.String(
		t,
		"Renamed Person",
		after[0].Attributes[constant.NameAttribute][0],
	)
	assert.FatalOnError(t, c.Delete(name))
	gone, g := c.Search("(uid=written)", []string{constant.NameAttribute})
	assert.FatalOnError(t, g)
	assert.Count(t, 0, gone)
}

func TestAddDuplicateConflicts(t *testing.T) {
	c := directory.NewEnvironment()
	name := distinguished("duplicate")
	assert.FatalOnError(t, c.Add(name, person("duplicate")))

	defer func() { assert.FatalOnError(t, c.Delete(name)) }()
	assert.True(t, conflict.Is(c.Add(name, person("duplicate"))))
}

func TestDeleteMissingNotFound(t *testing.T) {
	assert.True(
		t,
		not_found.Is(directory.NewEnvironment().Delete(distinguished("ghost"))),
	)
}

func TestModifyMissingNotFound(t *testing.T) {
	assert.True(
		t,
		not_found.Is(
			directory.NewEnvironment().Modify(
				distinguished("ghost"),
				map[string][]string{constant.NameAttribute: {"Nobody"}},
			),
		),
	)
}

func TestSetPasswordAuthenticates(t *testing.T) {
	c := directory.NewEnvironment()
	name := distinguished("rotated")
	attributes := person("rotated")
	attributes["userPassword"] = []string{"firstpassword"}
	assert.FatalOnError(t, c.Add(name, attributes))

	defer func() { assert.FatalOnError(t, c.Delete(name)) }()
	assert.FatalOnError(t, c.SetPassword(name, "secondpassword"))
	entry, e := c.Authenticate("rotated", "secondpassword")
	assert.FatalOnError(t, e)
	assert.String(t, "rotated", entry.Account)
}

func TestDeleteRemovesAuthentication(t *testing.T) {
	c := directory.NewEnvironment()
	name := distinguished("transient")
	attributes := person("transient")
	attributes["userPassword"] = []string{"transientpassword"}
	assert.FatalOnError(t, c.Add(name, attributes))
	_, e := c.Authenticate("transient", "transientpassword")
	assert.FatalOnError(t, e)
	assert.FatalOnError(t, c.Delete(name))
	_, f := c.Authenticate("transient", "transientpassword")
	assert.True(t, not_found.Is(f))
}
