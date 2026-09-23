package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	keepassConstant "github.com/funtimecoding/soil/pkg/keepass/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gocredentiald/constant"
	"path/filepath"
	"testing"
)

func TestListAndSearch(t *testing.T) {
	s := open(t, fixtureDatabase(t))
	assert.Count(t, 3, s.List())
	found := s.Search("forge")
	assert.Count(t, 1, found)
	assert.String(t, "Forge", found[0].Title)
	assert.String(t, "Root", found[0].Path)
}

func TestGetMasksPassword(t *testing.T) {
	s := open(t, fixtureDatabase(t))
	entry := s.Get(s.Search("forge")[0].Identifier)
	assert.NotNil(t, entry)
	assert.String(t, "•••", entry.Fields[constant.PasswordKey])
	assert.String(t, "•••", entry.Fields[constant.NotesKey])
	assert.String(t, "•••", entry.Fields["Extra"])
	assert.String(t, "alfa", entry.Fields[keepassConstant.UserNameKey])
	assert.String(t, "https://forge.example", entry.Fields[constant.LocatorKey])
}

func TestRevealedFieldList(t *testing.T) {
	path := fixtureDatabase(t)
	s := openRevealed(t, path, []string{"Extra"})
	entry := s.Get(s.Search("forge")[0].Identifier)
	assert.String(t, "extra-value", entry.Fields["Extra"])
	assert.String(t, "•••", entry.Fields[constant.NotesKey])
	assert.String(t, "•••", entry.Fields[constant.PasswordKey])
}

func TestRevealPassword(t *testing.T) {
	s := open(t, fixtureDatabase(t))
	password, found := s.Reveal(s.Search("forge")[0].Identifier)
	assert.True(t, found)
	assert.String(t, "hunter2", password)
}

func TestUpdatePersistsAcrossReopen(t *testing.T) {
	path := fixtureDatabase(t)
	s := open(t, path)
	identifier := s.Search("forge")[0].Identifier
	assert.FatalOnError(
		t,
		s.Update(
			identifier,
			map[string]string{constant.PasswordKey: "changed"},
		),
	)
	reopened := open(t, path)
	password, found := reopened.Reveal(identifier)
	assert.True(t, found)
	assert.String(t, "changed", password)
	backups, e := filepath.Glob(join.Empty(path, ".*.bak"))
	assert.FatalOnError(t, e)
	assert.Count(t, 1, backups)
}

func TestCreateAndDelete(t *testing.T) {
	path := fixtureDatabase(t)
	s := open(t, path)
	identifier, e := s.Create(
		"Root",
		"Fresh",
		map[string]string{
			keepassConstant.UserNameKey: "new",
			constant.PasswordKey:        "p",
		},
	)
	assert.FatalOnError(t, e)
	assert.Count(t, 4, s.List())
	assert.FatalOnError(t, s.Delete(identifier))
	assert.Count(t, 3, s.List())
}

func TestMove(t *testing.T) {
	path := fixtureDatabase(t)
	s := open(t, path)
	identifier := s.Search("forge")[0].Identifier
	assert.FatalOnError(t, s.Move(identifier, "Root/Environment"))
	assert.String(t, "Root/Environment", s.Search("forge")[0].Path)
}

func TestAudit(t *testing.T) {
	s := open(t, fixtureDatabase(t))
	report := s.Audit(3)
	assert.Count(t, 1, report.Stale)
	assert.String(t, "Ancient", report.Stale[0].Title)
}

func TestLoadGroup(t *testing.T) {
	s := open(t, fixtureDatabase(t))
	group := s.LoadGroup("example")
	assert.NotNil(t, group)
	assert.String(t, "example.test", group["EXAMPLE_HOST"])
	assert.String(t, "token-value", group["EXAMPLE_TOKEN"])
	assert.Count(t, 2, group)
	assert.Nil(t, s.LoadGroup("missing"))
}
