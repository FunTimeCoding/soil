package packages

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/strings/upper"
	"gitlab.com/gitlab-org/api/client-go/v2"
	"testing"
)

func TestLatest(t *testing.T) {
	assert.Any(
		t,
		map[string]*gitlab.Package{
			"Alfa":  {Name: "Alfa", Version: "v1.0.2"},
			"Bravo": {Name: "Bravo", Version: "v1.1.0"},
		},
		Latest(
			[]*gitlab.Package{
				{Name: upper.Alfa, Version: "v1.0.0"},
				{Name: upper.Alfa, Version: "v1.0.2"},
				{Name: upper.Alfa, Version: "v1.0.1"},
				{Name: upper.Bravo, Version: "v1.1.0"},
			},
		),
	)
}
