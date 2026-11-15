package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclean/clean/lab"
	"github.com/funtimecoding/soil/pkg/tool/goclean/clean/option"
	"github.com/funtimecoding/soil/pkg/tool/goclean/unit/clean_tester"
	"testing"
)

func TestPipelineCleanSparesRunningAndQueued(t *testing.T) {
	c := clean_tester.Project(t)
	lab.Pipeline(t.Context(), option.New(), c, clean_tester.Target())
	assert.Any(t, []int64{3}, c.DeletedPipelines())
}

func TestPipelineCleanAllDeletesRunningAndQueued(t *testing.T) {
	c := clean_tester.Project(t)
	o := option.New()
	o.All = true
	lab.Pipeline(t.Context(), o, c, clean_tester.Target())
	assert.Any(t, []int64{3, 4, 5}, c.DeletedPipelines())
}
