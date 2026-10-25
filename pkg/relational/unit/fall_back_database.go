package unit

import (
	"database/sql"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/relational/constant"
	"github.com/funtimecoding/soil/pkg/relational/lite"
	stringConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"time"
)

func fallBackDatabase() *sql.DB {
	m := lite.NewMemory()
	d, e := m.DB()
	errors.PanicOnError(e)
	_, e = d.Exec(constant.FixtureCreateEvent)
	errors.PanicOnError(e)
	_, e = d.Exec(
		constant.FixtureInsertEvent,
		stringConstant.LowerAlfa,
		time.Date(2026, 10, 25, 2, 30, 0, 0, time.FixedZone("", 2*60*60)),
	)
	errors.PanicOnError(e)
	_, e = d.Exec(
		constant.FixtureInsertEvent,
		stringConstant.LowerCharlie,
		time.Date(2026, 10, 25, 2, 30, 0, 0, time.FixedZone("", 60*60)),
	)
	errors.PanicOnError(e)

	return d
}
