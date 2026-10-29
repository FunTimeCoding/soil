package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/directory/mock_client"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/tool/gogated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogated/unit/base"
	"testing"
)

func TestAllowsWithoutDirectory(t *testing.T) {
	assert.True(t, allows(t, base.New(t), "anything"))
}

func TestAllowsLocalUser(t *testing.T) {
	d := mock_client.New("")
	b := base.NewWith(t, d)
	row, e := b.Store.UserByMail(constant.FixtureMail)
	assert.FatalOnError(t, e)
	assert.True(t, allows(t, b, row.Identifier))
}

func TestAllowsDirectoryMember(t *testing.T) {
	d := mock_client.New("")
	d.AddEntry("unique-1", "member", "member@example.test", "secret", true)
	b := base.NewWith(t, d)
	row, e := b.Service.AuthenticateUser(
		"member@example.test",
		"secret",
		constant.SourceDirectory,
	)
	assert.FatalOnError(t, e)
	assert.True(t, allows(t, b, row.Identifier))
}

func TestRevokesAfterGroupRemoval(t *testing.T) {
	d := mock_client.New("")
	d.AddEntry("unique-1", "member", "member@example.test", "secret", true)
	b := base.NewWith(t, d)
	row, e := b.Service.AuthenticateUser(
		"member@example.test",
		"secret",
		constant.SourceDirectory,
	)
	assert.FatalOnError(t, e)
	d.SetMember("member", false)
	assert.False(t, allows(t, b, row.Identifier))
}

func TestAnUnreachableDirectoryCannotConfirmMembership(t *testing.T) {
	d := mock_client.New("")
	d.AddEntry("unique-1", "member", "member@example.test", "secret", true)
	b := base.NewWith(t, d)
	row, e := b.Service.AuthenticateUser(
		"member@example.test",
		"secret",
		constant.SourceDirectory,
	)
	assert.FatalOnError(t, e)
	d.Failure = directoryDown()
	allowed, f := b.Service.DirectoryAllows(row.Identifier)
	assert.False(t, allowed)
	assert.True(t, unavailable.Is(f))
}

func TestAllowsUnknownIdentifierRefused(t *testing.T) {
	d := mock_client.New("")
	b := base.NewWith(t, d)
	assert.False(t, allows(t, b, "missing"))
}
