package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/directory/mock_client"
	"github.com/funtimecoding/soil/pkg/errors/conflict"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/tool/gogated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogated/unit/base"
	"testing"
)

func TestLocalUserNeverReachesDirectory(t *testing.T) {
	d := mock_client.New("")
	b := base.NewWith(t, d)
	d.Failure = directoryDown()
	row, e := b.Service.AuthenticateUser(
		constant.FixtureMail,
		constant.FixturePassword,
		constant.SourceLocal,
	)
	assert.FatalOnError(t, e)
	assert.NotNil(t, row)
	assert.String(t, "local", row.Source)
}

func TestDirectoryUserShadowCreated(t *testing.T) {
	d := mock_client.New("")
	b := base.NewWith(t, d)
	d.AddEntry("unique-1", "member", "member@example.test", "secret", true)
	row, e := b.Service.AuthenticateUser(
		"member@example.test",
		"secret",
		constant.SourceDirectory,
	)
	assert.FatalOnError(t, e)
	assert.NotNil(t, row)
	assert.String(t, "directory", row.Source)
	assert.String(t, "member", row.Account)
	assert.String(t, "unique-1", *row.DirectoryIdentifier)
}

func TestDirectoryUserByAccount(t *testing.T) {
	d := mock_client.New("")
	b := base.NewWith(t, d)
	d.AddEntry("unique-1", "member", "member@example.test", "secret", true)
	row, e := b.Service.AuthenticateUser(
		"member",
		"secret",
		constant.SourceDirectory,
	)
	assert.FatalOnError(t, e)
	assert.NotNil(t, row)
	assert.String(t, "member", row.Account)
}

func TestDirectoryUserReusesRow(t *testing.T) {
	d := mock_client.New("")
	b := base.NewWith(t, d)
	d.AddEntry("unique-1", "member", "member@example.test", "secret", true)
	first, e := b.Service.AuthenticateUser(
		"member@example.test",
		"secret",
		constant.SourceDirectory,
	)
	assert.FatalOnError(t, e)
	second, f := b.Service.AuthenticateUser(
		"member@example.test",
		"secret",
		constant.SourceDirectory,
	)
	assert.FatalOnError(t, f)
	assert.String(t, first.Identifier, second.Identifier)
}

func TestDirectoryMailChangeKeepsIdentity(t *testing.T) {
	d := mock_client.New("")
	b := base.NewWith(t, d)
	d.AddEntry("unique-1", "member", "member@example.test", "secret", true)
	first, e := b.Service.AuthenticateUser(
		"member@example.test",
		"secret",
		constant.SourceDirectory,
	)
	assert.FatalOnError(t, e)
	d.SetMail("unique-1", "renamed@example.test")
	second, f := b.Service.AuthenticateUser(
		"renamed@example.test",
		"secret",
		constant.SourceDirectory,
	)
	assert.FatalOnError(t, f)
	assert.String(t, first.Identifier, second.Identifier)
	assert.String(t, "renamed@example.test", second.Mail)
}

func TestDirectoryUserOutsideGroupRefused(t *testing.T) {
	d := mock_client.New("")
	b := base.NewWith(t, d)
	d.AddEntry("unique-2", "outsider", "outsider@example.test", "secret", false)
	row, e := b.Service.AuthenticateUser(
		"outsider@example.test",
		"secret",
		constant.SourceDirectory,
	)
	assert.FatalOnError(t, e)

	if row != nil {
		t.Errorf("expected refusal for a non-member")
	}
}

func TestDirectoryWrongPasswordRefused(t *testing.T) {
	d := mock_client.New("")
	b := base.NewWith(t, d)
	d.AddEntry("unique-1", "member", "member@example.test", "secret", true)
	row, e := b.Service.AuthenticateUser(
		"member@example.test",
		"wrong",
		constant.SourceDirectory,
	)
	assert.FatalOnError(t, e)

	if row != nil {
		t.Errorf("expected refusal for a wrong password")
	}
}

func TestLocalSourceRejectsDirectoryUser(t *testing.T) {
	d := mock_client.New("")
	b := base.NewWith(t, d)
	d.AddEntry("unique-1", "member", "member@example.test", "secret", true)
	row, e := b.Service.AuthenticateUser(
		"member@example.test",
		"secret",
		constant.SourceLocal,
	)
	assert.FatalOnError(t, e)

	if row != nil {
		t.Errorf("expected refusal for a directory user choosing local")
	}
}

func TestDirectorySourceRejectsLocalUser(t *testing.T) {
	d := mock_client.New("")
	b := base.NewWith(t, d)
	row, e := b.Service.AuthenticateUser(
		constant.FixtureMail,
		constant.FixturePassword,
		constant.SourceDirectory,
	)
	assert.FatalOnError(t, e)

	if row != nil {
		t.Errorf("expected refusal for a local user choosing directory")
	}
}

func TestDirectoryCollisionWithLocalMail(t *testing.T) {
	d := mock_client.New("")
	b := base.NewWith(t, d)
	d.AddEntry("unique-3", "collider", constant.FixtureMail, "secret", true)
	_, e := b.Service.AuthenticateUser(
		"collider",
		"secret",
		constant.SourceDirectory,
	)
	assert.True(t, conflict.Is(e))
}

func TestDirectoryUnreachableRefuses(t *testing.T) {
	d := mock_client.New("")
	b := base.NewWith(t, d)
	d.AddEntry("unique-1", "member", "member@example.test", "secret", true)
	d.Failure = directoryDown()
	_, e := b.Service.AuthenticateUser(
		"member@example.test",
		"secret",
		constant.SourceDirectory,
	)
	assert.True(t, unavailable.Is(e))
}

func TestWithoutDirectoryIgnoresSource(t *testing.T) {
	b := base.New(t)
	row, e := b.Service.AuthenticateUser(
		constant.FixtureMail,
		constant.FixturePassword,
		constant.SourceDirectory,
	)
	assert.FatalOnError(t, e)
	assert.NotNil(t, row)
	assert.String(t, "local", row.Source)
}
