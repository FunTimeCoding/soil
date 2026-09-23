package unit

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/keepass/constant"
	"github.com/funtimecoding/soil/pkg/system"
	constant1 "github.com/funtimecoding/soil/pkg/tool/gocredentiald/constant"
	"github.com/tobischo/gokeepasslib/v3"
	"github.com/tobischo/gokeepasslib/v3/wrappers"
	"path/filepath"
	"testing"
	"time"
)

func fixtureDatabase(t *testing.T) string {
	t.Helper()
	path := filepath.Join(t.TempDir(), "test.kdbx")
	database := gokeepasslib.NewDatabase(
		gokeepasslib.WithDatabaseKDBXVersion40(),
	)
	database.Credentials = gokeepasslib.NewPasswordCredentials("secret")
	root := gokeepasslib.NewGroup()
	root.Name = "Root"
	forge := gokeepasslib.NewEntry()
	forge.Values = append(
		forge.Values,
		value(constant1.TitleKey, "Forge", false),
		value(constant.UserNameKey, "alfa", false),
		value(constant1.PasswordKey, "hunter2", true),
		value(constant1.LocatorKey, "https://forge.example", false),
		value(constant1.NotesKey, "note text", false),
		value("Extra", "extra-value", false),
	)
	stale := gokeepasslib.NewEntry()
	stale.Values = append(
		stale.Values,
		value(constant1.TitleKey, "Ancient", false),
	)
	old := wrappers.Now()
	old.Time = time.Date(2018, 1, 1, 0, 0, 0, 0, time.UTC)
	stale.Times.LastModificationTime = &old
	root.Entries = append(root.Entries, forge, stale)
	environment := gokeepasslib.NewGroup()
	environment.Name = "Environment"
	group := gokeepasslib.NewEntry()
	group.Values = append(
		group.Values,
		value(constant1.TitleKey, "example", false),
		value("EXAMPLE_HOST", "example.test", false),
		value("EXAMPLE_TOKEN", "token-value", true),
	)
	environment.Entries = append(environment.Entries, group)
	root.Groups = append(root.Groups, environment)
	database.Content.Root.Groups = []gokeepasslib.Group{root}
	errors.PanicOnError(database.LockProtectedEntries())
	f := system.Create(path)
	errors.PanicOnError(gokeepasslib.NewEncoder(f).Encode(database))
	errors.PanicClose(f)

	return path
}
