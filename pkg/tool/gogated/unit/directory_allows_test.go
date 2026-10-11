package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/unreachable"
	"github.com/funtimecoding/soil/pkg/tool/gogated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogated/mock_directory"
	"github.com/funtimecoding/soil/pkg/tool/gogated/unit/base"
	"testing"
)

func TestAllowsWithoutDirectory(t *testing.T) {
	b := base.New(t)
	assert.True(t, b.Service.DirectoryAllows("anything"))
}

func TestAllowsLocalUser(t *testing.T) {
	b := base.New(t)
	d := mock_directory.New()
	b.Service.WithDirectory(d)
	row, e := b.Store.UserByMail(constant.FixtureMail)
	assert.FatalOnError(t, e)
	assert.True(t, b.Service.DirectoryAllows(row.Identifier))
}

func TestAllowsDirectoryMember(t *testing.T) {
	b := base.New(t)
	d := mock_directory.New()
	d.Add("unique-1", "member", "member@example.test", "secret", true)
	b.Service.WithDirectory(d)
	row, e := b.Service.AuthenticateUser(
		"member@example.test",
		"secret",
		constant.SourceDirectory,
	)
	assert.FatalOnError(t, e)
	assert.True(t, b.Service.DirectoryAllows(row.Identifier))
}

func TestRevokesAfterGroupRemoval(t *testing.T) {
	b := base.New(t)
	d := mock_directory.New()
	d.Add("unique-1", "member", "member@example.test", "secret", true)
	b.Service.WithDirectory(d)
	row, e := b.Service.AuthenticateUser(
		"member@example.test",
		"secret",
		constant.SourceDirectory,
	)
	assert.FatalOnError(t, e)
	d.SetMember("member", false)
	assert.False(t, b.Service.DirectoryAllows(row.Identifier))
}

func TestRevokesWhenDirectoryUnreachable(t *testing.T) {
	b := base.New(t)
	d := mock_directory.New()
	d.Add("unique-1", "member", "member@example.test", "secret", true)
	b.Service.WithDirectory(d)
	row, e := b.Service.AuthenticateUser(
		"member@example.test",
		"secret",
		constant.SourceDirectory,
	)
	assert.FatalOnError(t, e)
	d.SetFailure(unreachable.Format("directory down"))
	assert.False(t, b.Service.DirectoryAllows(row.Identifier))
}

func TestAllowsUnknownIdentifierRefused(t *testing.T) {
	b := base.New(t)
	d := mock_directory.New()
	b.Service.WithDirectory(d)
	assert.False(t, b.Service.DirectoryAllows("missing"))
}
