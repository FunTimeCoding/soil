package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/console/description"
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"testing"
)

func TestDescription(t *testing.T) {
	assert.Any(
		t,
		&description.Description{Title: "Alfa", Short: "Bravo"},
		description.New(stringsConstant.UpperAlfa, stringsConstant.UpperBravo),
	)
	assert.Any(
		t,
		&description.Description{
			Type:  "Alfa",
			Title: "Bravo",
			Short: "Charlie",
		},
		description.NewType(
			stringsConstant.UpperAlfa,
			stringsConstant.UpperBravo,
			stringsConstant.UpperCharlie,
		),
	)
}
