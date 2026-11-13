package unit

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/relational"
	"github.com/funtimecoding/soil/pkg/relational/constant"
	"github.com/funtimecoding/soil/pkg/relational/lite"
	"gorm.io/gorm"
	"testing"
)

func notes(t *testing.T) *gorm.DB {
	t.Helper()
	d := lite.NewMemory()
	t.Cleanup(func() { relational.CloseMapper(d) })
	errors.PanicOnError(d.Exec(constant.FixtureCreateNotes).Error)
	errors.PanicOnError(d.Exec(constant.FixtureNoteRow).Error)

	return d.Table(constant.FixtureOldTable)
}
