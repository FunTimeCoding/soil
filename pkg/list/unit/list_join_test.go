package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/list"
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"testing"
)

func TestJoin(t *testing.T) {
	assert.Any(
		t,
		[]string{"Alfa", "Bravo"},
		list.Join(
			[][]string{
				{stringsConstant.UpperAlfa},
				{stringsConstant.UpperBravo},
			},
		),
	)
}
