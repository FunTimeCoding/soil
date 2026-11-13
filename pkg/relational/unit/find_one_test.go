package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/relational"
	"github.com/funtimecoding/soil/pkg/relational/constant"
	"testing"
)

func TestFindOneFillsRow(t *testing.T) {
	row := map[string]any{}
	found, e := relational.FindOne(notes(t).Where("identifier = ?", 1), &row)
	assert.FatalOnError(t, e)
	assert.True(t, found)
	assert.Any(t, "kept", row[constant.FixtureOldColumn])
}

func TestFindOneMissIsNoError(t *testing.T) {
	row := map[string]any{}
	found, e := relational.FindOne(notes(t).Where("identifier = ?", 2), &row)
	assert.FatalOnError(t, e)
	assert.False(t, found)
}

func TestFindOneAnswersQueryError(t *testing.T) {
	row := map[string]any{}
	_, e := relational.FindOne(notes(t).Where("missing = ?", 1), &row)
	assert.Error(t, e)
}

func TestFetchOneFillsRow(t *testing.T) {
	row := map[string]any{}
	assert.FatalOnError(
		t,
		relational.FetchOne(
			notes(t).Where("identifier = ?", 1),
			&row,
			"note",
			1,
		),
	)
	assert.Any(t, "kept", row[constant.FixtureOldColumn])
}

func TestFetchOneMissIsNotFound(t *testing.T) {
	row := map[string]any{}
	e := relational.FetchOne(
		notes(t).Where("identifier = ?", 2),
		&row,
		"note",
		2,
	)
	assert.True(t, not_found.Is(e))
	assert.String(t, "note not found: 2", e.Error())
}
