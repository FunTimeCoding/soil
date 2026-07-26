package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"strings"
	"testing"
)

func TestMoveCommentTargetDrift(t *testing.T) {
	d := testutil.PrepareTestPackage(t, serviceTestdata("move-comment/src"))
	s := testService()
	r, e := s.MoveSymbols(
		d,
		"example/pkg/target",
		nil,
		"pkg/target/constant.go",
		"example/pkg/target/constant",
		"constant.go",
		false,
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	moved := readFixtureFile(t, d, "pkg/target/constant/constant.go")
	assertFormatted(t, moved)
	assert.StringContains(t, "const Host = \"bravo\" // host trailing", moved)
	assert.StringContains(t, "// Alpha document.", moved)
	assert.StringContains(t, "const Alpha = \"alfa\" // alpha trailing", moved)
	assert.True(
		t,
		strings.Index(moved, "// host trailing") <
			strings.Index(moved, "// Alpha document."),
	)
}

func TestMoveCommentSourceScrub(t *testing.T) {
	d := testutil.PrepareTestPackage(t, serviceTestdata("move-comment/src"))
	s := testService()
	r, e := s.MoveSymbols(
		d,
		"example/pkg/keeper",
		[]string{"moved"},
		"",
		"example/pkg/target/constant",
		"constant.go",
		false,
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	source := readFixtureFile(t, d, "pkg/keeper/constant.go")
	assertFormatted(t, source)
	assert.StringContains(t, "// Kept document.", source)
	assert.StringContains(t, "const kept = \"golf\" // kept trailing", source)
	assert.False(t, strings.Contains(source, "// Moved document."))
	assert.False(t, strings.Contains(source, "// moved trailing"))
	moved := readFixtureFile(t, d, "pkg/target/constant/constant.go")
	assertFormatted(t, moved)
	assert.StringContains(t, "// Moved document.", moved)
	assert.StringContains(t, "const Moved = \"hotel\" // moved trailing", moved)
}

func TestMoveCommentGroupScrub(t *testing.T) {
	d := testutil.PrepareTestPackage(t, serviceTestdata("move-comment/src"))
	s := testService()
	r, e := s.MoveSymbols(
		d,
		"example/pkg/keeper",
		[]string{"delta"},
		"",
		"example/pkg/target/constant",
		"constant.go",
		false,
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	source := readFixtureFile(t, d, "pkg/keeper/constant.go")
	assertFormatted(t, source)
	assert.StringContains(t, "epsilon = \"juliet\"", source)
	assert.False(t, strings.Contains(source, "// Delta document."))
	assert.False(t, strings.Contains(source, "// delta trailing"))
	moved := readFixtureFile(t, d, "pkg/target/constant/constant.go")
	assertFormatted(t, moved)
	assert.StringContains(t, "// Delta document.", moved)
	assert.StringContains(t, "const Delta = \"india\" // delta trailing", moved)
}
