package job

import (
	"github.com/funtimecoding/soil/pkg/gitlab/constant"
	"slices"
	"strings"
)

func (j *Job) Validate() {
	if j.Fail() && !slices.Contains(j.concern, constant.JobFailConcern) {
		j.concern = append(j.concern, constant.JobFailConcern)
	}

	if j.Trace != "" && !slices.Contains(j.concern, constant.JobTimeout) {
		if strings.Contains(j.Trace, constant.TraceTimeoutMatch) {
			j.concern = append(j.concern, constant.JobTimeout)
		}
	}
}
