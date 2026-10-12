package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/gitlab/packages"
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"gitlab.com/gitlab-org/api/client-go/v3"
	"testing"
)

func TestFindVersionOrLatest(t *testing.T) {
	assert.Any(
		t,
		&gitlab.Package{Name: "Alfa", Version: "v1.0.2"},
		packages.FindVersionOrLatest(
			[]*gitlab.Package{
				{Name: stringsConstant.UpperAlfa, Version: "v1.0.0"},
				{Name: stringsConstant.UpperAlfa, Version: "v1.0.2"},
				{Name: stringsConstant.UpperAlfa, Version: "v1.0.1"},
			},
			soil.LatestVersion,
		),
	)
	assert.Any(
		t,
		&gitlab.Package{Name: "Alfa", Version: "v1.0.1"},
		packages.FindVersionOrLatest(
			[]*gitlab.Package{
				{Name: stringsConstant.UpperAlfa, Version: "v1.0.0"},
				{Name: stringsConstant.UpperAlfa, Version: "v1.0.2"},
				{Name: stringsConstant.UpperAlfa, Version: "v1.0.1"},
			},
			"v1.0.1",
		),
	)
}
