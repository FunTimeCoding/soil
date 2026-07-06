package build_tag

import (
	"bufio"
	"github.com/funtimecoding/soil/pkg/errors"
	"go/build/constraint"
	"os"
	"strings"
)

func extract(path string) []string {
	file, e := os.Open(path)

	if e != nil {
		return nil
	}

	defer errors.PanicClose(file)
	var result []string
	s := bufio.NewScanner(file)

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
