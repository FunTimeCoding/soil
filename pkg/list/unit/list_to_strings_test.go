package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/list"
	"github.com/funtimecoding/soil/pkg/list/fixture"
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"testing"
)

func TestToStrings(t *testing.T) {
	assert.Strings(
		t,
		[]string{"Alfa", "Bravo"},
		list.ToStrings([]fixture.Text{constant.UpperAlfa, constant.UpperBravo}),
	)
}
