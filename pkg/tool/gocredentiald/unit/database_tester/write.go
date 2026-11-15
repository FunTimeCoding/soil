package database_tester

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/keepass"
	keepassConstant "github.com/funtimecoding/soil/pkg/keepass/constant"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/tool/gocredentiald/constant"
	"github.com/tobischo/gokeepasslib/v3"
	"github.com/tobischo/gokeepasslib/v3/wrappers"
	"path/filepath"
	"testing"
	"time"
)

func Write(t *testing.T) string {
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
		keepass.NewValue(constant.TitleKey, "Forge", false),
		keepass.NewValue(keepassConstant.UserNameKey, "alfa", false),
		keepass.NewValue(constant.PasswordKey, "hunter2", true),
		keepass.NewValue(constant.LocatorKey, "https://forge.example", false),
		keepass.NewValue(constant.NotesKey, "note text", false),
		keepass.NewValue("Extra", "extra-value", false),
	)
	stale := gokeepasslib.NewEntry()
	stale.Values = append(
		stale.Values,
		keepass.NewValue(constant.TitleKey, "Ancient", false),
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
		keepass.NewValue(constant.TitleKey, "example", false),
		keepass.NewValue("EXAMPLE_HOST", "example.test", false),
		keepass.NewValue("EXAMPLE_TOKEN", "token-value", true),
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
