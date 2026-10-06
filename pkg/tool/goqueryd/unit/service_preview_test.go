package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/generative/ollama"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/mock_reranker"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/service"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/unit/store_tester"
	"strings"
	"testing"
)

func TestPreviewReportsLinesTokensAndPieces(t *testing.T) {
	r := mock_reranker.New()
	r.SetAllowance(40)
	p := service.New(
		store_tester.OpenTestStore(t),
		ollama.New(),
		r,
	).Preview("results.md", constant.FixtureTableDocument)
	assert.Integer(t, 40, p.Allowance)
	assert.Count(t, 5, p.Chunks)
	assert.False(t, p.Chunks[0].Piece)
	assert.True(t, p.Chunks[2].Piece)
	assert.Integer(t, 11, p.Chunks[2].FirstLine)
	assert.Integer(t, 12, p.Chunks[2].LastLine)
}

func TestPreviewNamesTheLinesWhereAHeadingFixesAWindow(t *testing.T) {
	c := windowPreview(t, 560, strings.Repeat(constant.FixtureWindowLine, 200))
	assert.Integer(t, 600, c.Tokens)
	assert.True(t, c.CutChecked)
	assert.Integer(t, 6, c.CutLevel)
	assert.Count(t, 19, c.CutLines)
	assert.Integer(t, 95, c.CutLines[0])
	assert.Integer(t, 113, c.CutLines[18])
}

func TestPreviewMarksAWindowNoHeadingCanFix(t *testing.T) {
	c := windowPreview(t, 400, strings.Repeat(constant.FixtureWindowLine, 200))
	assert.True(t, c.CutChecked)
	assert.Integer(t, 0, c.CutLevel)
	assert.Count(t, 0, c.CutLines)
}

func TestPreviewAccountsForALaterHeadingThatWinsTheCut(t *testing.T) {
	c := windowPreview(
		t,
		560,
		join.Empty(
			strings.Repeat(constant.FixtureWindowLine, 117),
			constant.FixtureWindowHeading,
			strings.Repeat(constant.FixtureWindowLine, 82),
		),
	)
	assert.Integer(t, 585, c.Tokens)
	assert.Integer(t, 2, c.CutLevel)
	assert.Integers(t, []int{111, 112, 113}, c.CutLines)
}
