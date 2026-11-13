package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/web/layout"
	"strings"
	"testing"
	"time"
)

func TestOptionalTimeCellShowsTimeWhenSet(t *testing.T) {
	var b strings.Builder
	at := time.Date(2025, time.March, 4, 5, 6, 0, 0, time.Local)
	assert.FatalOnError(t, layout.OptionalTimeCell(&at, "never").Render(&b))
	assert.String(
		t,
		`<td class="time-cell"><small>2025-03-04 05:06</small></td>`,
		b.String(),
	)
}

func TestOptionalTimeCellNamesAbsenceInSharedLook(t *testing.T) {
	var b strings.Builder
	assert.FatalOnError(t, layout.OptionalTimeCell(nil, "never").Render(&b))
	assert.String(
		t,
		`<td class="time-cell"><small class="absent">never</small></td>`,
		b.String(),
	)
}

func TestClampCellKeepsWholeTextInPage(t *testing.T) {
	var b strings.Builder
	assert.FatalOnError(t, layout.ClampCell("alfa bravo").Render(&b))
	assert.String(
		t,
		`<td class="clamp-cell"><span data-tooltip="alfa bravo">alfa bravo</span></td>`,
		b.String(),
	)
}

func TestOptionalTimeCellWithoutWordStaysEmpty(t *testing.T) {
	var b strings.Builder
	assert.FatalOnError(t, layout.OptionalTimeCell(nil, "").Render(&b))
	assert.String(t, `<td class="time-cell"></td>`, b.String())
}
