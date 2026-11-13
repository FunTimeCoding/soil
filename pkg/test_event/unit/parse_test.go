package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/test_event"
	"github.com/funtimecoding/soil/pkg/test_event/constant"
	"testing"
)

func TestParseSkipsLinesOutsideStream(t *testing.T) {
	events := test_event.Parse(constant.FixtureMixedStream)
	assert.Count(t, 10, events)
	assert.String(t, "start", events[0].Action)
}

func TestParseReadsBuildLines(t *testing.T) {
	events := test_event.Parse(constant.FixtureBuildFailureStream)
	assert.String(t, "probe/broken [probe/broken.test]", events[0].ImportPath)
	assert.String(t, "probe/broken [probe/broken.test]", events[5].FailedBuild)
}
