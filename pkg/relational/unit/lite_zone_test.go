package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/relational/constant"
	"github.com/funtimecoding/soil/pkg/relational/lite/connection"
	"github.com/funtimecoding/soil/pkg/relational/lite/zone"
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"testing"
	"time"
)

func TestZoneStoresUniversalTime(t *testing.T) {
	d := fallBackDatabase()
	defer errors.PanicClose(d)
	assert.Strings(
		t,
		[]string{"charlie", "alfa"},
		textColumn(d, constant.FixtureEventNewestFirst),
	)
	assert.Strings(
		t,
		[]string{"2026-10-25 00:30:00+00:00", "2026-10-25 01:30:00+00:00"},
		textColumn(d, constant.FixtureEventStoredTime),
	)
}

func TestZoneComparesUniversalArgument(t *testing.T) {
	d := fallBackDatabase()
	defer errors.PanicClose(d)
	assert.Strings(
		t,
		[]string{"charlie"},
		textColumn(
			d,
			constant.FixtureEventAfter,
			time.Date(2026, 10, 25, 2, 0, 0, 0, time.FixedZone("", 60*60)),
		),
	)
}

func TestZoneReadsLocalTime(t *testing.T) {
	d := fallBackDatabase()
	defer errors.PanicClose(d)
	var read time.Time
	errors.PanicOnError(
		d.QueryRow(
			constant.FixtureEventTime,
			stringsConstant.LowerAlfa,
		).Scan(&read),
	)
	assert.True(t, read.Location() == time.Local)
	assert.True(t, read.Equal(time.Date(2026, 10, 25, 0, 30, 0, 0, time.UTC)))
}

func TestNormalizeConvertsStoredOffsets(t *testing.T) {
	d := connection.NewMemory()
	defer errors.PanicClose(d)
	_, e := d.Exec(constant.FixtureCreateEvent)
	errors.PanicOnError(e)

	for _, v := range [][]any{
		{
			stringsConstant.LowerAlfa,
			"2026-10-25 02:30:00.123456+02:00",
			"2026-13-45 02:30:00+02:00",
		},
		{stringsConstant.LowerBravo, "2026-10-25 00:30:00+00:00", nil},
		{stringsConstant.LowerCharlie, "2026-10-25 02:30:00+01:00", nil},
	} {
		_, f := d.Exec(constant.FixtureInsertNotedEvent, v...)
		errors.PanicOnError(f)
	}

	zone.Normalize(d)
	assert.Strings(
		t,
		[]string{
			"2026-10-25 00:30:00.123456+00:00",
			"2026-10-25 00:30:00+00:00",
			"2026-10-25 01:30:00+00:00",
		},
		textColumn(d, constant.FixtureEventStoredTime),
	)
	assert.Strings(
		t,
		[]string{"2026-13-45 02:30:00+02:00"},
		textColumn(d, constant.FixtureEventNote),
	)
}

func TestNormalizeRunsOnce(t *testing.T) {
	d := connection.NewMemory()
	defer errors.PanicClose(d)
	_, e := d.Exec(constant.FixtureCreateEvent)
	errors.PanicOnError(e)
	zone.Normalize(d)
	_, e = d.Exec(
		constant.FixtureInsertEvent,
		stringsConstant.LowerAlfa,
		"2026-10-25 02:30:00+02:00",
	)
	errors.PanicOnError(e)
	zone.Normalize(d)
	assert.Strings(
		t,
		[]string{"2026-10-25 02:30:00+02:00"},
		textColumn(d, constant.FixtureEventStoredTime),
	)
}
