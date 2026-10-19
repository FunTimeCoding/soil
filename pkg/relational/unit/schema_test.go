package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/relational/constant"
	"github.com/funtimecoding/soil/pkg/relational/lite"
	"github.com/funtimecoding/soil/pkg/relational/schema"
	"testing"
)

func TestARenamedColumnKeepsItsValues(t *testing.T) {
	d := lite.NewMemory()
	errors.PanicOnError(d.Exec(constant.FixtureCreateNotes).Error)
	errors.PanicOnError(d.Exec(constant.FixtureNoteRow).Error)
	renames := map[string]string{
		constant.FixtureOldColumn: constant.FixtureNewColumn,
	}
	schema.RenameColumns(d, constant.FixtureOldTable, renames)
	schema.RenameColumns(d, constant.FixtureOldTable, renames)
	var content string
	errors.PanicOnError(
		d.Raw(constant.FixtureRenamedContent).Scan(&content).Error,
	)
	assert.String(t, "kept", content)
	assert.False(
		t,
		d.Migrator().HasColumn(
			constant.FixtureOldTable,
			constant.FixtureOldColumn,
		),
	)
}

func TestARenameOntoAnExistingColumnFails(t *testing.T) {
	d := lite.NewMemory()
	errors.PanicOnError(d.Exec(constant.FixtureCreateBoth).Error)
	assert.Panic(
		t,
		func() {
			schema.RenameColumns(
				d,
				constant.FixtureOldTable,
				map[string]string{
					constant.FixtureOldColumn: constant.FixtureNewColumn,
				},
			)
		},
	)
}

func TestARenamedTableKeepsItsRows(t *testing.T) {
	d := lite.NewMemory()
	errors.PanicOnError(d.Exec(constant.FixtureCreateNotes).Error)
	errors.PanicOnError(d.Exec(constant.FixtureNoteRow).Error)
	schema.RenameTable(d, constant.FixtureOldTable, constant.FixtureNewTable)
	schema.RenameTable(d, constant.FixtureOldTable, constant.FixtureNewTable)
	var body string
	errors.PanicOnError(d.Raw(constant.FixtureMovedBody).Scan(&body).Error)
	assert.String(t, "kept", body)
	assert.False(t, d.Migrator().HasTable(constant.FixtureOldTable))
}

func TestADroppedColumnAndTableAreGone(t *testing.T) {
	d := lite.NewMemory()
	errors.PanicOnError(d.Exec(constant.FixtureCreateNotes).Error)
	schema.DropColumns(d, constant.FixtureOldTable, constant.FixtureStaleColumn)
	schema.DropColumns(d, constant.FixtureOldTable, constant.FixtureStaleColumn)
	assert.False(
		t,
		d.Migrator().HasColumn(
			constant.FixtureOldTable,
			constant.FixtureStaleColumn,
		),
	)
	assert.True(
		t,
		d.Migrator().HasColumn(
			constant.FixtureOldTable,
			constant.FixtureOldColumn,
		),
	)
	schema.DropTable(d, constant.FixtureOldTable)
	schema.DropTable(d, constant.FixtureOldTable)
	assert.False(t, d.Migrator().HasTable(constant.FixtureOldTable))
}
