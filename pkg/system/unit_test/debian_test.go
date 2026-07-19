package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/constant"
	system "github.com/funtimecoding/soil/pkg/system/constant"
	"github.com/funtimecoding/soil/pkg/system/debian"
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
