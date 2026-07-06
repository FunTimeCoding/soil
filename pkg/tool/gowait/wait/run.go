package wait

import "github.com/funtimecoding/soil/pkg/tool/gowait/wait/option"

func Run(o *option.Wait) {
	if o.Locator != "" {
		Locator(o)
	} else if o.File != "" {
		File(o)
	} else if o.Process != "" {
		Process(o)
	}
}
