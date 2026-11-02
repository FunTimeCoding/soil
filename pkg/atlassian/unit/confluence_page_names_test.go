package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/page"
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"testing"
)

func TestConfluencePageNames(t *testing.T) {
	assert.Strings(
		t,
		[]string{"Alfa", "Bravo"},
		page.Names(
			[]*page.Page{
				{Name: stringsConstant.UpperAlfa},
				{Name: stringsConstant.UpperBravo},
			},
		),
	)
}
