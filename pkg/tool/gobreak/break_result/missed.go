package break_result

import "github.com/funtimecoding/soil/pkg/tool/gobreak/constant"

func (v *Result) Missed() bool {
	switch v.Expect {
	case constant.StatusKilled:
		return v.Status != constant.StatusKilled &&
			v.Status != constant.StatusPanic
	case constant.StatusSurvived:
		return v.Status != constant.StatusSurvived &&
			v.Status != constant.StatusNoCoverage
	}

	return false
}
