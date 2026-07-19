package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/gitlab/packages"
	"github.com/funtimecoding/soil/pkg/strings/upper"
	"gitlab.com/gitlab-org/api/client-go/v2"
	"testing"
)

func TestFindVersion(t *testing.T) {
	assert.Any(
		t,
		&gitlab.Package{Name: "Alfa", Version: "v1.0.1"},
		packages.FindVersion(
			[]*gitlab.Package{
				{Name: upper.Alfa, Version: "v1.0.0"},
				{Name: upper.Alfa, Version: "v1.0.2"},
				{Name: upper.Alfa, Version: "v1.0.1"},
			},
			"v1.0.1",
		),
	)
}
