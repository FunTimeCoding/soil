package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/console/description"
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"testing"
)

func TestDescription(t *testing.T) {
	assert.Any(
		t,
		&description.Description{Title: "Alfa", Short: "Bravo"},
		description.New(constant.UpperAlfa, constant.UpperBravo),
	)
	assert.Any(
		t,
		&description.Description{
			Type:  "Alfa",
			Title: "Bravo",
			Short: "Charlie",
		},
		description.NewType(
			constant.UpperAlfa,
			constant.UpperBravo,
			constant.UpperCharlie,
		),
	)
}
