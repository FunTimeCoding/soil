package runner

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/system/run"
	"github.com/funtimecoding/soil/pkg/tool/goagentd/constant"
)

func (r *Runner) execute(intent string) {
	command := run.New()
	command.Directory = r.workspace
	command.NoPanic()
	command.Start("claude", "--print", "-p", intent)
	r.mutex.Lock()
	r.state = constant.RunnerDone

	if command.Error != nil {
		r.result = fmt.Sprintf(
			"%s\n\nerror: %s",
			command.OutputString,
			command.Error,
		)
	} else {
		r.result = command.OutputString
	}

	r.mutex.Unlock()
}
