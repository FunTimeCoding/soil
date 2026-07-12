package git

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/git/constant"
	"github.com/funtimecoding/soil/pkg/system"
	"testing"
	"time"
)

func TestCommitTimes(t *testing.T) {
	times, e := CommitTimes(system.ParentDirectory(constant.Depth))
	assert.FatalOnError(t, e)
	modified, okay := times["go.mod"]
	assert.True(t, okay)
	assert.True(t, modified.After(time.Unix(0, 0)))
}

func TestCommitTimesOutsideRepository(t *testing.T) {
	_, e := CommitTimes(t.TempDir())
	assert.Error(t, e)
}

func TestUncommittedFiles(t *testing.T) {
	_, e := UncommittedFiles(system.ParentDirectory(constant.Depth))
	assert.FatalOnError(t, e)
}
