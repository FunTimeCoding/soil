package runner

import (
	"github.com/funtimecoding/soil/pkg/tool/goagentd/constant"
	"sync"
)

type Runner struct {
	mutex     sync.Mutex
	state     constant.RunnerState
	result    string
	workspace string
}
