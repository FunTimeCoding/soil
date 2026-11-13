package build_tag

import (
	"bufio"
	"go/build/constraint"
	"io"
	"strings"
)

func extractFrom(r io.Reader) []string {
	var result []string
	s := bufio.NewScanner(r)

	for s.Scan() {
		line := s.Text()

		if line == "" || strings.HasPrefix(line, "//") {
			if constraint.IsGoBuild(line) {
				expression, f := constraint.Parse(line)

				if f != nil {
					continue
				}

				collect(expression, &result)
			}

			continue
		}

		if strings.HasPrefix(line, "package ") {
			break
		}
	}

	return result
}
