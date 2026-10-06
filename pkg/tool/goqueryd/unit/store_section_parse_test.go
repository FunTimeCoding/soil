package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store/section"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store/section/parser"
	"testing"
)

func TestParseSplitsAtHeadingsOutsideFences(t *testing.T) {
	sections := parser.Parse(constant.FixtureSectionDocument)
	assert.Count(t, 4, sections)
	assert.String(t, "", sections[0].Title)
	assert.String(t, "# Alfa", sections[1].Title)
	assert.Integer(t, 1, sections[1].Level)
	assert.String(t, "## Bravo", sections[2].Title)
	assert.Integer(t, 9, sections[2].FirstLine)
	assert.Integer(t, 17, sections[2].LastLine)
	assert.String(t, "## Charlie", sections[3].Title)
}

func TestParseClassifiesBlocks(t *testing.T) {
	sections := parser.Parse(constant.FixtureSectionDocument)
	assert.String(t, "front matter", sections[0].Blocks[0].Kind)
	assert.String(t, "heading", sections[1].Blocks[0].Kind)
	assert.String(t, "paragraph", sections[1].Blocks[1].Kind)
	assert.Integer(t, 6, sections[1].Blocks[1].FirstLine)
	assert.Integer(t, 7, sections[1].Blocks[1].LastLine)
	assert.String(t, "list", sections[2].Blocks[1].Kind)
	assert.String(t, "code", sections[2].Blocks[2].Kind)
	assert.Integer(t, 16, sections[2].Blocks[2].LastLine)
	assert.String(t, "table", sections[3].Blocks[1].Kind)
}

func TestWordsSkipHeadingsAndMarkup(t *testing.T) {
	assert.Strings(
		t,
		[]string{"Intro", "line", "one.", "Intro", "line", "two."},
		section.Words(parser.Parse(constant.FixtureSectionDocument)[1]),
	)
}
