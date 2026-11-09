package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/relational/constant"
	"github.com/funtimecoding/soil/pkg/relational/lite/connection"
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"path/filepath"
	"testing"
	"time"
)

func TestReadThenWriteTransactionKeepsItsWrite(t *testing.T) {
	d := connection.New(
		logger.New(t.Context()),
		filepath.Join(t.TempDir(), constant.FixtureLiteFile),
	)
	defer errors.PanicClose(d)
	_, e := d.Exec(constant.FixtureCreateEvent)
	assert.FatalOnError(t, e)
	first, e := d.BeginTx(t.Context(), nil)
	assert.FatalOnError(t, e)
	var count int
	assert.FatalOnError(
		t,
		first.QueryRow(constant.FixtureEventCount).Scan(&count),
	)
	other := make(chan error, 1)
	go func() {
		_, f := d.Exec(
			constant.FixtureInsertEvent,
			stringsConstant.LowerCharlie,
			time.Now(),
		)
		other <- f
	}()
	time.Sleep(constant.FixtureOtherWriterStart)
	_, e = first.Exec(
		constant.FixtureInsertEvent,
		stringsConstant.LowerAlfa,
		time.Now(),
	)
	assert.FatalOnError(t, e)
	assert.FatalOnError(t, first.Commit())

	select {
	case f := <-other:
		assert.FatalOnError(t, f)
	case <-time.After(constant.FixtureOtherWriterWait):
		t.Fatal("the other writer never finished")
	}

	assert.FatalOnError(t, d.QueryRow(constant.FixtureEventCount).Scan(&count))
	assert.Integer(t, 2, count)
}
