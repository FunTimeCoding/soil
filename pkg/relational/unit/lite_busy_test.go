package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/relational/constant"
	"github.com/funtimecoding/soil/pkg/relational/lite"
	"github.com/funtimecoding/soil/pkg/relational/lite/connection"
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"path/filepath"
	"testing"
	"time"
)

func TestAStaleReadSnapshotWritingIsBusy(t *testing.T) {
	d := connection.New(
		logger.New(t.Context()),
		filepath.Join(t.TempDir(), constant.FixtureLiteFile),
	)
	defer errors.PanicClose(d)
	_, e := d.Exec(constant.FixtureCreateEvent)
	assert.FatalOnError(t, e)
	reader, e := d.Conn(t.Context())
	assert.FatalOnError(t, e)
	defer errors.PanicClose(reader)
	_, e = reader.ExecContext(t.Context(), constant.FixtureBeginDeferred)
	assert.FatalOnError(t, e)
	defer func() {
		_, f := reader.ExecContext(t.Context(), constant.FixtureRollback)
		assert.FatalOnError(t, f)
	}()
	var count int
	assert.FatalOnError(
		t,
		reader.QueryRowContext(
			t.Context(),
			constant.FixtureEventCount,
		).Scan(&count),
	)
	_, e = d.Exec(
		constant.FixtureInsertEvent,
		stringsConstant.LowerCharlie,
		time.Now(),
	)
	assert.FatalOnError(t, e)
	_, f := reader.ExecContext(
		t.Context(),
		constant.FixtureInsertEvent,
		stringsConstant.LowerAlfa,
		time.Now(),
	)
	assert.True(t, lite.IsBusy(f))
}
