package wait

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/system/run"
	"github.com/funtimecoding/soil/pkg/tool/gowait/wait/option"
	"time"
)

func Process(o *option.Wait) {
	start := time.Now()
	attempt := 0

	for {
		attempt++

		if o.Verbose {
			fmt.Printf("Check %d\n", attempt)
		}

		r := run.New().NoPanic()
		r.Start("pgrep", "-f", o.Process)

		if r.Error != nil {
			fmt.Printf(
				"Done after %v\n",
				time.Since(start).Round(time.Second),
			)

			return
		}

		if o.Verbose {
			fmt.Printf(
				"Running: %s",
				run.New().Start("pgrep", "-fa", o.Process),
			)
		}

		if time.Since(start) > o.Timeout {
			panic(fmt.Sprintf("timeout after %v", o.Timeout))
		}

		time.Sleep(interval)
	}
}
