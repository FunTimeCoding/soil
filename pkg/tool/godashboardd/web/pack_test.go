package web

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/assert/fixture"
	"github.com/funtimecoding/go-library/pkg/system/constant"
	"github.com/funtimecoding/go-library/pkg/tool/godashboardd/board"
	"testing"
)

func TestPack(t *testing.T) {
	b := board.Load(fixture.Path(constant.BoardPath, "pack.yaml"))
	packed := pack(b.Tail.Sections, b.Tail.Columns)
	assert.Integer(t, 2, len(packed))
	assert.Integer(t, 1, len(packed[0]))
	assert.String(t, "five", packed[0][0].Name)
	assert.Integer(t, 3, len(packed[1]))
	assert.String(t, "two", packed[1][0].Name)
}

func TestPackDeterministic(t *testing.T) {
	b := board.Load(fixture.Path(constant.BoardPath, "pack.yaml"))
	first := pack(b.Tail.Sections, b.Tail.Columns)
	second := pack(b.Tail.Sections, b.Tail.Columns)
	assert.Integer(t, len(first[0]), len(second[0]))
	assert.String(t, first[0][0].Name, second[0][0].Name)
	assert.String(t, first[1][0].Name, second[1][0].Name)
}
