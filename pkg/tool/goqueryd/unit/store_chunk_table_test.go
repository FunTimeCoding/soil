package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"testing"
)

func TestTableOverAllowanceSplitsBetweenRows(t *testing.T) {
	chunks := tableChunks(40)
	assert.Count(t, 5, chunks)
	assert.String(
		t,
		"### Table 1\nTable 1. Alfa bravo charlie.\n(part 2 of 3, rows 3–4 of 6)\n\n| a | b |\n|---|---|\n| 3 a | 3 b |\n| 4 a | 4 b |\n",
		chunks[2].Text,
	)
	assert.String(t, "\nClosing paragraph.\n", chunks[4].Text)
}

func TestTablePiecesCoverTheSourceContiguously(t *testing.T) {
	chunks := tableChunks(40)

	for i := 1; i < 4; i++ {
		assert.Integer(
			t,
			chunks[i+1].Position,
			chunks[i].Position+chunks[i].Length,
		)
	}

	assert.String(
		t,
		"| a | b |\n|---|---|\n| 1 a | 1 b |\n| 2 a | 2 b |\n",
		constant.FixtureTableDocument[chunks[1].Position:chunks[1].Position+chunks[1].Length],
	)
}

func TestTableThatTipsItsWindowOverLeavesTheText(t *testing.T) {
	chunks := tableChunks(50)
	assert.Count(t, 4, chunks)
	assert.String(
		t,
		"# Results\n\n### Table 1\n\nTable 1. Alfa bravo charlie.\n\n",
		chunks[0].Text,
	)
	assert.StringContains(t, "(part 1 of 2, rows 1–4 of 6)", chunks[1].Text)
	assert.StringContains(t, "(part 2 of 2, rows 5–6 of 6)", chunks[2].Text)
}

func TestTableWithinAllowanceStaysInTheText(t *testing.T) {
	chunks := tableChunks(100)
	assert.Count(t, 1, chunks)
	assert.String(
		t,
		"# Results\n\n### Table 1\n\nTable 1. Alfa bravo charlie.\n\n| a | b |\n|---|---|\n| 1 a | 1 b |\n| 2 a | 2 b |\n| 3 a | 3 b |\n| 4 a | 4 b |\n| 5 a | 5 b |\n| 6 a | 6 b |\n\nClosing paragraph.\n",
		chunks[0].Text,
	)
}
