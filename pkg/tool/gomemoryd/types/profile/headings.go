package profile

import (
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"slices"
	"strings"
)

func Headings(
	text string,
	section string,
) []*Heading {
	var result []*Heading
	inside := false

	for _, line := range strings.Split(text, stringsConstant.Unix) {
		if slices.Contains(constant.SectionHeadings, line) {
			inside = line == section

			continue
		}

		if !inside {
			continue
		}

		if h := headingLine(line); h != nil {
			result = append(result, h)
		}
	}

	return result
}
