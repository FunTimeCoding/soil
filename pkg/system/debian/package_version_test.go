package debian

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/constant"
	system "github.com/funtimecoding/soil/pkg/system/constant"
	"testing"
)

func TestPackageVersion(t *testing.T) {
	assert.String(
		t,
		"example_1.0.0-1_amd64",
		PackageVersion(
			"example",
			constant.DefaultVersion,
			1,
			system.AMD64,
		),
	)
}
