package pipeline_view

import (
	"github.com/funtimecoding/soil/pkg/gitlab/job"
	"github.com/funtimecoding/soil/pkg/gitlab/pipeline_detail"
)

type View struct {
	Instance string
	Project  int64
	Pipeline int64
	Detail   *pipeline_detail.Detail
	Jobs     []*job.Job
	Selected *job.Job
	Trace    string
}
