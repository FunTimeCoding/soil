package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/space"
	"github.com/funtimecoding/soil/pkg/strings/upper"
	"testing"
)

func TestConfluenceSpaceNames(t *testing.T) {
	assert.Strings(
		t,
		[]string{"Alfa", "Bravo"},
		space.Names([]*space.Space{{Name: upper.Alfa}, {Name: upper.Bravo}}),
	)
}
