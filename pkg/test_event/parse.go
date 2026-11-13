package test_event

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/strings/split"
	"github.com/funtimecoding/soil/pkg/test_event/constant"
	"strings"
)

func Parse(output string) []*Event {
	var result []*Event

	for _, line := range split.NewLine(output) {
		if !strings.HasPrefix(line, constant.ObjectOpening) {
			continue
		}

		e := &Event{}
		errors.PanicOnError(notation.Decode(line, e))
		result = append(result, e)
	}

	return result
}
