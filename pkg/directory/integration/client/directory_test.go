//go:build ci

package client

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/directory"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	bootstrap(directory.NewEnvironment())
	os.Exit(m.Run())
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
