package unit

import (
	"bufio"
	"github.com/funtimecoding/soil/pkg/tool/gomonitord/constant"
	"strings"
	"testing"
	"time"
)

func readEvent(
	t *testing.T,
	r *bufio.Reader,
) string {
	t.Helper()
	result := make(chan string, 1)
	go func() {
		for {
			line, e := r.ReadString('\n')

			if e != nil {
				result <- ""

				return
			}

			if payload, found := strings.CutPrefix(
				line,
				constant.EventPrefix,
			); found {
				result <- strings.TrimSpace(payload)

				return
			}
		}
	}()

	select {
	case v := <-result:
		return v
	case <-time.After(5 * time.Second):
		t.Fatal("no event within 5 seconds")

		return ""
	}
}
