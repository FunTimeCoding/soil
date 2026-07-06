package constant

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestConstant(t *testing.T) {
	assert.String(t, "control", ControlFile)
	assert.String(t, "DEBIAN", PackageConfigurationDirectory)
	assert.String(t, ".deb", PackageExtension)
}
