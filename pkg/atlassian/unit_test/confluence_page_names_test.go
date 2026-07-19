package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/page"
	"github.com/funtimecoding/soil/pkg/strings/upper"
	"testing"
)

func TestConfluencePageNames(t *testing.T) {
	assert.Strings(
		t,
		[]string{"Alfa", "Bravo"},
		page.Names([]*page.Page{{Name: upper.Alfa}, {Name: upper.Bravo}}),
	)
}
