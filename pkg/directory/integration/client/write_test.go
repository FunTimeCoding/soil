//go:build ci

package client

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/directory"
	"github.com/funtimecoding/soil/pkg/directory/constant"
	"github.com/funtimecoding/soil/pkg/errors/conflict"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"testing"
)

func person(account string) map[string][]string {
	return map[string][]string{
		"objectClass":             {"inetOrgPerson"},
		constant.AccountAttribute: {account},
		constant.NameAttribute:    {account},
		"sn":                      {account},
	}
}

func distinguished(account string) string {
	return fmt.Sprintf("uid=%s,ou=people,dc=example,dc=test", account)
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
