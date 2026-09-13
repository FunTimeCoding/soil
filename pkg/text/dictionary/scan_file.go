package dictionary

import (
	"bufio"
	"github.com/funtimecoding/soil/pkg/errors"
	library "github.com/funtimecoding/soil/pkg/strings"
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/system"
	"strings"
)

func ScanFile(
	path string,
	w map[string]*WordUsage,
) {
	f := system.Open(path)
	defer errors.LogClose(f)
	b := strings.Builder{}
	s := bufio.NewScanner(f)

	for s.Scan() {
		b.WriteString(s.Text())
		b.WriteString(constant.Unix)
	}

	errors.PanicOnError(s.Err())
	raw := b.String()
	token := map[string]bool{}
	tokenize(strings.ToLower(raw), token)
	tokenize(
		strings.ToLower(
			splitDigit(strings.ReplaceAll(library.SplitCase(raw), "_", " ")),
		),
		token,
	)

	for _, u := range w {
		if !u.Used && token[u.lower] {
			u.Used = true
		}
	}
}
