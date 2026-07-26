package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/page"
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"testing"
)

func TestConfluencePageNames(t *testing.T) {
	assert.Strings(
		t,
		[]string{"Alfa", "Bravo"},
		page.Names([]*page.Page{{Name: constant.UpperAlfa}, {Name: constant.UpperBravo}}),
	)
}
