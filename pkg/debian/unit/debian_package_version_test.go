package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/debian"
	system "github.com/funtimecoding/soil/pkg/system/constant"
	"testing"
)

func TestPackageVersion(t *testing.T) {
	assert.String(
		t,
		"example_1.0.0-1_amd64",
		debian.PackageVersion(
			"example",
			constant.DefaultVersion,
			1,
			system.AMD64,
		),
	)
}
