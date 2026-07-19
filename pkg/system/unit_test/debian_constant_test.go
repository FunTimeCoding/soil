package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/system/debian/constant"
	"testing"
)

func TestDebianConstant(t *testing.T) {
	assert.String(t, "control", constant.ControlFile)
	assert.String(t, "DEBIAN", constant.PackageConfigurationDirectory)
	assert.String(t, ".deb", constant.PackageExtension)
}
