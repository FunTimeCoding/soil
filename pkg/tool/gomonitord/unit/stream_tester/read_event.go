package stream_tester

import (
	"bufio"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gomonitord/constant"
	"strings"
	"testing"
)

func ReadEvent(
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

	return assert.Signalled(t, result)
}
