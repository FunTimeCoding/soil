package base

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"net/http"
	"testing"
	"time"
)

func waitReady(
	t *testing.T,
	port int,
) {
	t.Helper()
	locator := fmt.Sprintf("http://localhost:%d/json/version", port)
	deadline := time.Now().Add(20 * time.Second)

	for time.Now().Before(deadline) {
		response, e := http.Get(locator)

		if e == nil {
			errors.LogClose(response.Body)

			if response.StatusCode == http.StatusOK {
				return
			}
		}

		time.Sleep(100 * time.Millisecond)
	}

	t.Fatalf("debugging endpoint on port %d never became ready", port)
}
