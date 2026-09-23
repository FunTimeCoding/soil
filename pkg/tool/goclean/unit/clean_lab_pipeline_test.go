package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclean/clean/lab"
	"github.com/funtimecoding/soil/pkg/tool/goclean/clean/option"
	"testing"
)

func TestPipelineCleanSparesRunningAndQueued(t *testing.T) {
	c := seedProject()
	lab.Pipeline(option.New(), c, seedTarget())
	assert.Any(t, []int64{3}, c.DeletedPipelines())
}

func TestPipelineCleanAllDeletesRunningAndQueued(t *testing.T) {
	c := seedProject()
	o := option.New()
	o.All = true
	lab.Pipeline(o, c, seedTarget())
	assert.Any(t, []int64{3, 4, 5}, c.DeletedPipelines())
}
