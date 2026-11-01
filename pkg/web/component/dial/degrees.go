package dial

import "github.com/funtimecoding/soil/pkg/web/constant"

func degrees(share float64) float64 {
	return constant.DialDegrees * min(1, max(0, share))
}
