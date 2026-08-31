package web

import (
	"fmt"
	argument "github.com/funtimecoding/soil/pkg/argument/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/constant"
)

func detailLink(project int64, pipeline int64, job int64) string {
	result := fmt.Sprintf(
		"%s?%s=%d&%s=%d",
		constant.PipelinePath,
		argument.Project,
		project,
		argument.Pipeline,
		pipeline,
	)

	if job != 0 {
		result = fmt.Sprintf("%s&%s=%d", result, argument.Job, job)
	}

	return result
}
