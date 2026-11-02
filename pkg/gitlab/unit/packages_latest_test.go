package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/gitlab/packages"
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"gitlab.com/gitlab-org/api/client-go/v3"
	"testing"
)

func TestPackagesLatest(t *testing.T) {
	assert.Any(
		t,
		map[string]*gitlab.Package{
			"Alfa":  {Name: "Alfa", Version: "v1.0.2"},
			"Bravo": {Name: "Bravo", Version: "v1.1.0"},
		},
		packages.Latest(
			[]*gitlab.Package{
				{Name: stringsConstant.UpperAlfa, Version: "v1.0.0"},
				{Name: stringsConstant.UpperAlfa, Version: "v1.0.2"},
				{Name: stringsConstant.UpperAlfa, Version: "v1.0.1"},
				{Name: stringsConstant.UpperBravo, Version: "v1.1.0"},
			},
		),
	)
}
