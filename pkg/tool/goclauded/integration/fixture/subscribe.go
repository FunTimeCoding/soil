package fixture

import (
	"bufio"
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/integration/base"
	"net/http"
	"strings"
	"testing"
)

func Subscribe(
	t *testing.T,
	s *base.Server,
) (<-chan Event, func()) {
	t.Helper()
	l := fmt.Sprintf(
		"http://localhost:%d/event?subscribe=roster,activity,summary",
		s.Port,
	)
	r, e := http.Get(l)
	assert.FatalOnError(t, e)
	events := make(chan Event, 10)
	go func() {
		scanner := bufio.NewScanner(r.Body)
		var name string
		var lines []string

		for scanner.Scan() {
			line := scanner.Text()

			if strings.HasPrefix(line, "event: ") {
				name = strings.TrimPrefix(line, "event: ")
				lines = nil
			} else if strings.HasPrefix(line, "data: ") {
				lines = append(lines, strings.TrimPrefix(line, "data: "))
			} else if line == "" && name != "" {
				events <- Event{Name: name, Payload: join.NewLine(lines)}
				name = ""
				lines = nil
			}
		}

		close(events)
	}()

	return events, func() { errors.PanicClose(r.Body) }
}
