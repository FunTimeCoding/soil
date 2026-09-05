package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert"
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"testing"
)

func TestSortByAge(t *testing.T) {
	dayOne := assert.NewDay(1)
	dayTwo := assert.NewDay(2)
	descending := alert.SortByAge(
		[]*alert.Alert{
			{Name: constant.UpperAlfa, Start: new(dayOne)},
			{Name: constant.UpperBravo, Start: new(dayTwo)},
		},
		false,
	)
	assert.String(t, "Bravo", descending[0].Name)
	assert.String(t, "Alfa", descending[1].Name)
	ascending := alert.SortByAge(
		[]*alert.Alert{
			{Name: constant.UpperAlfa, Start: new(dayOne)},
			{Name: constant.UpperBravo, Start: new(dayTwo)},
		},
		true,
	)
	assert.String(t, "Alfa", ascending[0].Name)
	assert.String(t, "Bravo", ascending[1].Name)
}
