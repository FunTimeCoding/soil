package connector

import (
	"bufio"
	"context"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/stream_event"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"net/http"
	"strings"
)

func (c *Client) streamOnce(
	x context.Context,
	name string,
	kinds []string,
	from string,
	events chan<- *stream_event.Event,
	last *string,
) bool {
	r, e := http.NewRequestWithContext(
		x,
		http.MethodGet,
		c.streamLocator(name, kinds),
		nil,
	)

	if e != nil {
		return false
	}

	r.Header.Set(
		webConstant.Authorization,
		join.Space(webConstant.Bearer, c.token),
	)

	if from != "" {
		r.Header.Set(webConstant.LastEvent, from)
	}

	response, e := c.streamClient().Do(r)

	if e != nil {
		return false
	}

	defer errors.LogClose(response.Body)

	if response.StatusCode != http.StatusOK {
		return false
	}

	scanner := bufio.NewScanner(response.Body)
	scanner.Buffer(
		make([]byte, 0, constant.StreamLineBuffer),
		constant.StreamLineBuffer,
	)
	delivered := false
	identifier := ""
	var lines []string

	for scanner.Scan() {
		line := scanner.Text()

		switch {
		case strings.HasPrefix(line, constant.StreamIdentifierPrefix):
			identifier = strings.TrimPrefix(
				line,
				constant.StreamIdentifierPrefix,
			)
		case strings.HasPrefix(line, constant.StreamPayloadPrefix):
			lines = append(
				lines,
				strings.TrimPrefix(line, constant.StreamPayloadPrefix),
			)
		case line == "" && len(lines) > 0:
			result := stream_event.New()

			if notation.DecodeBytes([]byte(join.NewLine(lines)), result) == nil {
				select {
				case events <- result:
					delivered = true

					if identifier != "" {
						*last = identifier
					}
				case <-x.Done():
					return delivered
				}
			}

			identifier = ""
			lines = nil
		}
	}

	return delivered
}
