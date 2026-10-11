package fixture

import (
	"bufio"
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/base"
	"github.com/funtimecoding/soil/pkg/web"
	"net/http"
	"strings"
	"testing"
)

func SubscribeStream(
	t *testing.T,
	s *base.Server,
	name string,
	kinds string,
) (<-chan Event, func()) {
	t.Helper()
	l := fmt.Sprintf(
		"http://localhost:%d/event-stream?subscriber=%s&kinds=%s",
		s.Port,
		name,
		kinds,
	)
	r, e := http.NewRequest(http.MethodGet, l, nil)
	assert.FatalOnError(t, e)
	web.Bearer(r, generativeConstant.ModelContextTestToken)
	response, e := http.DefaultClient.Do(r)
	assert.FatalOnError(t, e)
	events := make(chan Event, 50)
	go func() {
		scanner := bufio.NewScanner(response.Body)
		var identifier string
		var name string
		var lines []string

		for scanner.Scan() {
			line := scanner.Text()

			switch {
			case strings.HasPrefix(line, constant.StreamIdentifierPrefix):
				identifier = strings.TrimPrefix(
					line,
					constant.StreamIdentifierPrefix,
				)
			case strings.HasPrefix(line, "event: "):
				name = strings.TrimPrefix(line, "event: ")
				lines = nil
			case strings.HasPrefix(line, constant.StreamPayloadPrefix):
				lines = append(
					lines,
					strings.TrimPrefix(line, constant.StreamPayloadPrefix),
				)
			case line == "" && name != "":
				events <- Event{
					Name:       name,
					Identifier: identifier,
					Payload:    join.NewLine(lines),
				}
				name = ""
				identifier = ""
				lines = nil
			}
		}

		close(events)
	}()

	return events, func() { errors.PanicClose(response.Body) }
}
