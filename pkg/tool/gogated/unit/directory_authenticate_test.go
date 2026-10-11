package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/conflict"
	"github.com/funtimecoding/soil/pkg/errors/unreachable"
	"github.com/funtimecoding/soil/pkg/tool/gogated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogated/mock_directory"
	"github.com/funtimecoding/soil/pkg/tool/gogated/unit/base"
	"testing"
)

func TestLocalUserNeverReachesDirectory(t *testing.T) {
	b := base.New(t)
	d := mock_directory.New()
	d.SetFailure(unreachable.Format("directory down"))
	b.Service.WithDirectory(d)
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
	assert.NotNil(t, row)
	assert.String(t, "directory", row.Source)
	assert.String(t, "member", row.Account)
	assert.String(t, "unique-1", *row.DirectoryIdentifier)
}

func TestDirectoryUserByAccount(t *testing.T) {
	b := base.New(t)
	d := mock_directory.New()
	d.Add("unique-1", "member", "member@example.test", "secret", true)
	b.Service.WithDirectory(d)
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
	b := base.New(t)
	d := mock_directory.New()
	d.Add("unique-1", "member", "member@example.test", "secret", true)
	b.Service.WithDirectory(d)
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
	b := base.New(t)
	d := mock_directory.New()
	d.Add("unique-1", "member", "member@example.test", "secret", true)
	b.Service.WithDirectory(d)
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
	b := base.New(t)
	d := mock_directory.New()
	d.Add("unique-2", "outsider", "outsider@example.test", "secret", false)
	b.Service.WithDirectory(d)
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
	b := base.New(t)
	d := mock_directory.New()
	d.Add("unique-1", "member", "member@example.test", "secret", true)
	b.Service.WithDirectory(d)
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
	b := base.New(t)
	d := mock_directory.New()
	d.Add("unique-1", "member", "member@example.test", "secret", true)
	b.Service.WithDirectory(d)
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
	b := base.New(t)
	d := mock_directory.New()
	b.Service.WithDirectory(d)
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
	b := base.New(t)
	d := mock_directory.New()
	d.Add("unique-3", "collider", constant.FixtureMail, "secret", true)
	b.Service.WithDirectory(d)
	_, e := b.Service.AuthenticateUser(
		"collider",
		"secret",
		constant.SourceDirectory,
	)
	assert.True(t, conflict.Is(e))
}

func TestDirectoryUnreachableRefuses(t *testing.T) {
	b := base.New(t)
	d := mock_directory.New()
	d.Add("unique-1", "member", "member@example.test", "secret", true)
	b.Service.WithDirectory(d)
	d.SetFailure(unreachable.Format("directory down"))
	_, e := b.Service.AuthenticateUser(
		"member@example.test",
		"secret",
		constant.SourceDirectory,
	)
	assert.True(t, unreachable.Is(e))
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
