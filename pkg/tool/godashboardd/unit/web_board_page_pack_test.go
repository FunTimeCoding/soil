package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/assert/fixture"
	"github.com/funtimecoding/soil/pkg/system/constant"
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/board"
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/web/board_page"
	"testing"
)

func TestPack(t *testing.T) {
	b := board.Load(fixture.Path(constant.BoardPath, "pack.yaml"))
	packed := board_page.Pack(b.Tail.Sections, b.Tail.Columns)
	assert.Integer(t, 2, len(packed))
	assert.Integer(t, 1, len(packed[0]))
	assert.String(t, "five", packed[0][0].Name)
	assert.Integer(t, 3, len(packed[1]))
	assert.String(t, "two", packed[1][0].Name)
}

func TestPackDeterministic(t *testing.T) {
	b := board.Load(fixture.Path(constant.BoardPath, "pack.yaml"))
	first := board_page.Pack(b.Tail.Sections, b.Tail.Columns)
	second := board_page.Pack(b.Tail.Sections, b.Tail.Columns)
	assert.Integer(t, len(first[0]), len(second[0]))
	assert.String(t, first[0][0].Name, second[0][0].Name)
	assert.String(t, first[1][0].Name, second[1][0].Name)
}
