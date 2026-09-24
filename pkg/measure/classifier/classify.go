package classifier

import (
	"github.com/funtimecoding/soil/pkg/measure/constant"
	"github.com/funtimecoding/soil/pkg/measure/count"
	"github.com/funtimecoding/soil/pkg/strings/split"
	"strings"
)

func (c *Classifier) Classify(content string) *count.Count {
	result := count.New()
	c.open = nil
	c.quote = ""
	c.raw = false
	content = strings.TrimPrefix(content, constant.ByteOrderMark)
	lines := split.NewLine(content)

	if len(lines) > 0 && lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	for i, line := range lines {
		trimmed := strings.TrimSpace(line)

		switch {
		case trimmed == "":
			result.Blank++
		case i == 0 && strings.HasPrefix(trimmed, constant.ShebangPrefix):
			result.Code++
		case c.line(trimmed):
			result.Comment++
		default:
			result.Code++
		}

		if !c.raw {
			c.quote = ""
		}
	}

	return result
}
