package run_result

import "github.com/funtimecoding/soil/pkg/tool/goagentd/constant"

func New(
	state constant.RunnerState,
	result string,
) *Result {
	return &Result{State: state, Result: result}
}
