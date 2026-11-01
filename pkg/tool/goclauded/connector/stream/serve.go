package stream

import (
	"bufio"
	"context"
	"github.com/funtimecoding/soil/pkg/errors/dropped"
	"github.com/funtimecoding/soil/pkg/errors/unreachable"
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/stream_event"
	"strings"
)

func (s *Stream) Serve(x context.Context) error {
	s.mutex.Lock()
	b := s.body
	s.mutex.Unlock()

	if b == nil {
		return unreachable.Format("stream not open")
	}

	scanner := bufio.NewScanner(b)
	scanner.Buffer(
		make([]byte, 0, constant.StreamLineBuffer),
		constant.StreamLineBuffer,
	)
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
				case s.events <- result:
					if identifier != "" {
						s.last = identifier
					}
				case <-x.Done():
					return nil
				}
			}

			identifier = ""
			lines = nil
		}
	}

	if e := scanner.Err(); e != nil {
		return e
	}

	return dropped.Format("%s", constant.StreamEnded)
}
