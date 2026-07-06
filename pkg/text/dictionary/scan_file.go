package dictionary

import (
	"bufio"
	"github.com/funtimecoding/soil/pkg/errors"
	library "github.com/funtimecoding/soil/pkg/strings"
	"github.com/funtimecoding/soil/pkg/strings/separator"
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
		b.WriteString(separator.Unix)
	}

	errors.PanicOnError(s.Err())
	content := strings.ToLower(b.String())

	for wordKey, u := range w {
		if !u.Used && library.HasWord(content, wordKey) {
			u.Used = true
		}

		// TODO: Some words are parts of other words, is there a better way?
		//  Some variables are capitalized, like noiseWithMatch and Match is the word, so no word boundary, but capital letter to indicate new word
		if !u.Used && strings.Contains(content, wordKey) {
			u.Used = true
		}
	}
}
