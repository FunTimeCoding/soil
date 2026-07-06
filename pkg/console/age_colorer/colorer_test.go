package age_colorer

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/console/age_colorer/age_fixture"
	"github.com/funtimecoding/soil/pkg/math/range_mapping"
	"github.com/funtimecoding/soil/pkg/math/ranges"
	"testing"
	"time"
)

func TestColorer(t *testing.T) {
	g := age_fixture.New(0 * time.Hour)
	y := age_fixture.New(15 * time.Hour)
	r := age_fixture.New(30 * time.Hour)
	c := Default(g, y, r)
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
		c.mapping,
	)
	c.Set(g)
	c.Set(y)
	c.Set(r)
	console.GreenInstance.EnableColor()
	console.YellowInstance.EnableColor()
	console.RedInstance.EnableColor()
	// Not sure if function pointers can be compared, so compare output
	assert.String(
		t,
		console.Green("%s", "g"),
		g.AgeColor()("g"),
	)
	assert.String(
		t,
		console.Yellow("%s", "y"),
		y.AgeColor()("y"),
	)
	assert.String(
		t,
		console.Red("%s", "r"),
		r.AgeColor()("r"),
	)
}
