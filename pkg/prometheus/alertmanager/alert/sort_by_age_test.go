package alert

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/strings/upper"
	"testing"
)

func TestSortByAge(t *testing.T) {
	dayOne := assert.NewDay(1)
	dayTwo := assert.NewDay(2)
	descending := SortByAge(
		[]*Alert{
			{Name: upper.Alfa, Start: new(dayOne)},
			{Name: upper.Bravo, Start: new(dayTwo)},
		},
		false,
	)
	assert.String(t, "Bravo", descending[0].Name)
	assert.String(t, "Alfa", descending[1].Name)
	ascending := SortByAge(
		[]*Alert{
			{Name: upper.Alfa, Start: new(dayOne)},
			{Name: upper.Bravo, Start: new(dayTwo)},
		},
		true,
	)
	assert.String(t, "Alfa", ascending[0].Name)
	assert.String(t, "Bravo", ascending[1].Name)
}
