package coverage

import (
	"github.com/funtimecoding/soil/pkg/crap/constant"
	"github.com/funtimecoding/soil/pkg/strings/slice"
	"github.com/funtimecoding/soil/pkg/strings/split"
	"github.com/funtimecoding/soil/pkg/system/run"
	"strings"
)

func ListTests(binary string) []string {
	r := run.New()
	var result []string
	output := r.Start(binary, constant.TestList, constant.AnyTest)

	for _, line := range slice.StripEmpty(
		split.NewLine(strings.TrimSpace(output)),
	) {
		if strings.HasPrefix(line, constant.TestPrefix) {
			result = append(result, line)
		}
	}

	return result
}
