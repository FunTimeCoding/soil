package dictionary

import (
	"bufio"
	"github.com/funtimecoding/soil/pkg/errors"
	soilStrings "github.com/funtimecoding/soil/pkg/strings"
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
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
		b.WriteString(stringsConstant.Unix)
	}

	errors.PanicOnError(s.Err())
	raw := b.String()
	token := map[string]bool{}
	tokenize(strings.ToLower(raw), token)
	tokenize(
		strings.ToLower(
			splitDigit(
				strings.ReplaceAll(
					soilStrings.SplitCase(raw),
					"_",
					stringsConstant.Space,
				),
			),
		),
		token,
	)

	for _, u := range w {
		if !u.Used && token[u.lower] {
			u.Used = true
		}
	}
}
