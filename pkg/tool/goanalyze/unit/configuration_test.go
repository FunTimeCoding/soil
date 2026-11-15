package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/system/system_tester"
	"github.com/funtimecoding/soil/pkg/tool/goanalyze/configuration"
	"github.com/funtimecoding/soil/pkg/tool/goanalyze/constant"
	"testing"
)

func TestConfigurationAbsentLeavesCommentOff(t *testing.T) {
	assert.False(t, configuration.Load(t.TempDir()).Comment)
}

func TestConfigurationTurnsCommentOn(t *testing.T) {
	root := t.TempDir()
	system_tester.Write(t, root, constant.ConfigurationPath, "comment: true\n")
	assert.True(t, configuration.Load(root).Comment)
}

func TestConfigurationMalformedPanics(t *testing.T) {
	root := t.TempDir()
	system_tester.Write(t, root, constant.ConfigurationPath, "comment: [\n")
	assert.Panic(t, func() { configuration.Load(root) })
}
