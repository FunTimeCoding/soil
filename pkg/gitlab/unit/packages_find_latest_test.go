package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/gitlab/packages"
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"gitlab.com/gitlab-org/api/client-go/v2"
	"testing"
)

func TestFindLatest(t *testing.T) {
	assert.Any(
		t,
		&gitlab.Package{Name: "Alfa", Version: "v1.0.2"},
		packages.FindLatest(
			[]*gitlab.Package{
				{Name: constant.UpperAlfa, Version: "v1.0.0"},
				{Name: constant.UpperAlfa, Version: "v1.0.2"},
				{Name: constant.UpperAlfa, Version: "v1.0.1"},
			},
		),
	)
}
