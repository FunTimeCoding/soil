package unit_test

import (
	"github.com/fatih/color"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/console/score_colorer"
	"github.com/funtimecoding/soil/pkg/console/score_colorer/score_fixture"
	"github.com/funtimecoding/soil/pkg/math/range_mapping"
	"github.com/funtimecoding/soil/pkg/math/ranges"
	"testing"
)

func TestScoreColorerDefault(t *testing.T) {
	g := score_fixture.New(0)
	y := score_fixture.New(15)
	r := score_fixture.New(30)
	c := score_colorer.Default(g, y, r)
	assert.Any(
		t,
		[]*range_mapping.Mapping{
			{
				Range: ranges.Range{
					L: 0,
					R: 0.3333333333333333,
				},
				Value: "green",
			},
			{
				Range: ranges.Range{
					L: 0.3333333333333333,
					R: 0.6666666666666666,
				},
				Value: "yellow",
			},
			{
				Range: ranges.Range{
					L: 0.6666666666666666,
					R: 1,
				},
				Value: "red",
			},
		},
		c.Mapping(),
	)
	c.Set(g)
	c.Set(y)
	c.Set(r)
	color.NoColor = false
	// Not sure if function pointers can be compared, so compare output
	assert.String(
		t,
		constant.Green("%s", "g"),
		g.ScoreColor()("g"),
	)
	assert.String(
		t,
		constant.Yellow("%s", "y"),
		y.ScoreColor()("y"),
	)
	assert.String(
		t,
		constant.Red("%s", "r"),
		r.ScoreColor()("r"),
	)
}

func TestZeroScore(t *testing.T) {
	a := score_fixture.New(0)
	b := score_fixture.New(0)
	c := score_colorer.Default(a, b)
	c.Set(a)
	c.Set(b)
	color.NoColor = false
	assert.String(t, constant.Green("%s", "a"), a.ScoreColor()("a"))
	assert.String(t, constant.Green("%s", "b"), b.ScoreColor()("b"))
}
