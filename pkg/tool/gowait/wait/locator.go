package wait

import (
	"context"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/tool/gowait/constant"
	"github.com/funtimecoding/soil/pkg/tool/gowait/wait/option"
	"io"
	"net/http"
	"strings"
	"time"
)

func Locator(o *option.Wait) {
	if o.Contains == "" {
		system.Exitf(1, "flag empty: contains\n")
	}

	x, cancel := context.WithTimeout(context.Background(), o.Timeout)
	defer cancel()
	c := &http.Client{Timeout: 10 * time.Second}
	attempt := 0

	for {
		attempt++

		if o.Verbose {
			console.Format("%s %d\n", o.Locator, attempt)
		}

		r, getFail := http.NewRequestWithContext(
			x,
			http.MethodGet,
			o.Locator,
			nil,
		)
		errors.PanicOnError(getFail)
		result, doFail := c.Do(r)

		if doFail != nil {
			if o.Verbose {
				console.Format("Do fail: %v\n", doFail)
			}
		} else {
			body, readFail := io.ReadAll(result.Body)
			errors.LogClose(result.Body)

			if readFail != nil {
				if o.Verbose {
					console.Format("Read fail: %v\n", readFail)
				}
			} else if result.StatusCode == http.StatusOK {
				content := string(body)

				if o.Verbose {
					console.Format("Response: %s\n", content)
				}

				if strings.Contains(content, o.Contains) {
					console.Line("Found")

					return
				}
			} else if o.Verbose {
				console.Format("Status: %d\n", result.StatusCode)
			}
		}

		select {
		case <-x.Done():
			panic("timeout")
		case <-time.After(constant.Interval):
			// pass
		}
	}
}
