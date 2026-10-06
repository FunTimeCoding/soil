package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestOversizeListsFilesOverTheAllowanceWorstFirst(t *testing.T) {
	r := oversizeReport(t, "")
	assert.String(t, "mock", r.Model)
	assert.Integer(t, 5, r.Allowance)
	assert.Count(t, 2, r.Files)
	assert.String(t, "bravo.md", r.Files[0].Path)
	assert.Integer(t, 10, r.Files[0].Worst)
	assert.String(t, "charlie.md", r.Files[1].Path)
	assert.Integer(t, 7, r.Files[1].Worst)
}

func TestOversizeChunkCarriesLinesBytesAndTokens(t *testing.T) {
	c := oversizeReport(t, "").Files[0].Chunks[0]
	assert.Integer(t, 0, c.Index)
	assert.Integer(t, 1, c.FirstLine)
	assert.Integer(t, 4, c.LastLine)
	assert.Integer(t, 49, c.Bytes)
	assert.Integer(t, 10, c.Tokens)
}

func TestOversizeCollectionFilterExcludesOtherCollections(t *testing.T) {
	assert.Count(t, 0, oversizeReport(t, "other").Files)
}
