package runner

import "github.com/funtimecoding/soil/pkg/tool/goagentd/constant"

func New(workspace string) *Runner {
	return &Runner{state: constant.RunnerIdle, workspace: workspace}
}
