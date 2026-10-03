package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goanalyze/configuration"
	"testing"
)

func TestConfigurationAbsentLeavesCommentOff(t *testing.T) {
	assert.False(t, configuration.Load(t.TempDir()).Comment)
}

func TestConfigurationTurnsCommentOn(t *testing.T) {
	root := t.TempDir()
	writeConfiguration(root, "comment: true\n")
	assert.True(t, configuration.Load(root).Comment)
}

func TestConfigurationMalformedPanics(t *testing.T) {
	defer func() { assert.NotNil(t, recover()) }()
	root := t.TempDir()
	writeConfiguration(root, "comment: [\n")
	configuration.Load(root)
}
