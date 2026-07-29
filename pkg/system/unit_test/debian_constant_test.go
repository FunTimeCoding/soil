package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/system/constant"
	"testing"
)

func TestDebianConstant(t *testing.T) {
	assert.String(t, "control", constant.DebianControlFile)
	assert.String(t, "DEBIAN", constant.DebianPackageConfigurationDirectory)
	assert.String(t, ".deb", constant.DebianPackageExtension)
}
