package web

import (
	"fmt"
	argumentConstant "github.com/funtimecoding/soil/pkg/argument/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/constant"
)

func detailLink(
	project int64,
	pipeline int64,
	job int64,
) string {
	result := fmt.Sprintf(
		"%s?%s=%d&%s=%d",
		constant.PipelinePath,
		argumentConstant.Project,
		project,
		argumentConstant.Pipeline,
		pipeline,
	)

	if job != 0 {
		result = fmt.Sprintf("%s&%s=%d", result, argumentConstant.Job, job)
	}

	return result
}
