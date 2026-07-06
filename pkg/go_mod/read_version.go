package go_mod

import (
	"bufio"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/system"
	"regexp"
	"strings"
)

func ReadVersion(mod string) string {
	f := system.Open(mod)
	defer errors.PanicClose(f)
	r := regexp.MustCompile(`^go\s+(\d+\.\d+(?:\.\d+)?)`)
	s := bufio.NewScanner(f)

	for s.Scan() {
		if m := r.FindStringSubmatch(strings.TrimSpace(s.Text())); m != nil {
			return m[1]
		}
	}

	errors.PanicOnError(s.Err())

	return ""
}
