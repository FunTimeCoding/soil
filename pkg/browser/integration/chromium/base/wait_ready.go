package base

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/time/wait_tester"
	"net/http"
	"testing"
	"time"
)

func waitReady(
	t *testing.T,
	c *http.Client,
	port int,
) {
	t.Helper()
	locator := fmt.Sprintf("http://localhost:%d/json/version", port)
	wait_tester.Until(
		t,
		func() bool {
			response, e := c.Get(locator)

			if e != nil {
				return false
			}

			errors.LogClose(response.Body)

			return response.StatusCode == http.StatusOK
		},
		20*time.Second,
	)
}
