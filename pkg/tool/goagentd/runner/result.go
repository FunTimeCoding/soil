package runner

import "github.com/funtimecoding/soil/pkg/tool/goagentd/constant"

type Result struct {
	State  constant.RunnerState `json:"state"`
	Result string               `json:"result,omitempty"`
}
