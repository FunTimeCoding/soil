package base

import (
	"bufio"
	"strings"
)

func (s *Server) AwaitLine(
	r *bufio.Reader,
	substring string,
) {
	s.t.Helper()

	for {
		line, e := r.ReadString('\n')

		if e != nil {
			s.t.Fatalf("stream ended before a line with %q: %v", substring, e)
		}

		if strings.Contains(line, substring) {
			return
		}
	}
}
