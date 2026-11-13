package break_result

import "github.com/funtimecoding/soil/pkg/tool/gobreak/constant"

func (v *Result) Missed() bool {
	switch v.Expect {
	case constant.ExpectRed:
		return v.Kind != constant.ResultRed &&
			v.Kind != constant.ResultPanic
	case constant.ExpectSurvive:
		return v.Kind != constant.ResultSurvived &&
			v.Kind != constant.ResultUnreached
	}

	return false
}
