package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/unit/oversize_tester"
	"math"
	"testing"
)

func TestOversizeListsFilesOverAllowanceWorstFirst(t *testing.T) {
	r := oversize_tester.Report(t, "")
	assert.String(t, "mock", r.Model)
	assert.Integer(t, 5, r.Allowance)
	assert.Count(t, 2, r.Files)
	assert.String(t, "bravo.md", r.Files[0].Path)
	assert.Integer(t, 10, r.Files[0].Worst)
	assert.String(t, "charlie.md", r.Files[1].Path)
	assert.Integer(t, 7, r.Files[1].Worst)
}

func TestOversizeChunkCarriesLinesBytesAndTokens(t *testing.T) {
	c := oversize_tester.Report(t, "").Files[0].Chunks[0]
	assert.Integer(t, 0, c.Index)
	assert.Integer(t, 1, c.FirstLine)
	assert.Integer(t, 4, c.LastLine)
	assert.Integer(t, 49, c.Bytes)
	assert.Integer(t, 10, c.Tokens)
}

func TestOversizeListsChunkOverEmbedderAlone(t *testing.T) {
	r := oversize_tester.ReportWithin(t, "", math.MaxInt, 8)
	assert.Integer(t, 8, r.EmbedAllowance)
	assert.Count(t, 1, r.Files)
	assert.String(t, "bravo.md", r.Files[0].Path)
	assert.Integer(t, 10, r.Files[0].Chunks[0].EmbedTokens)
}

func TestOversizeCollectionFilterExcludesOtherCollections(t *testing.T) {
	assert.Count(t, 0, oversize_tester.Report(t, "other").Files)
}
