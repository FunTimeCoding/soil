package digest

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/digest/event"
	"strings"
)

func enumerated(alias string, events []*event.Event) string {
	result := []string{
		fmt.Sprintf("[%s] %d new · %s", alias, len(events), span(events)),
	}

	for _, e := range events {
		result = append(result, join.Empty(constant.DigestIndent, line(e)))
	}

	return strings.Join(result, "\n")
}
