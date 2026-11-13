package break_result

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gobreak/constant"
)

func (v *Result) Line() string {
	detail := v.Detail

	if len(v.Tests) > 0 {
		detail = join.Space(v.Tests...)
	}

	marker := ""

	if v.Missed() {
		marker = fmt.Sprintf(constant.UnexpectedMarker, v.Expect)
	}

	return fmt.Sprintf(constant.LineFormat, v.Name, v.Kind, detail, marker)
}
