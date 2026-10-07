package pointer

import "github.com/funtimecoding/soil/pkg/lint/constant"

func NewResolution(verdict constant.Verdict) *Resolution {
	return &Resolution{Verdict: verdict}
}
