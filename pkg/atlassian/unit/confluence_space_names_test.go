package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/space"
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"testing"
)

func TestConfluenceSpaceNames(t *testing.T) {
	assert.Strings(
		t,
		[]string{"Alfa", "Bravo"},
		space.Names(
			[]*space.Space{
				{Name: stringsConstant.UpperAlfa},
				{Name: stringsConstant.UpperBravo},
			},
		),
	)
}
