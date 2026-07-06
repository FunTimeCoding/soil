package float

import (
	"github.com/funtimecoding/soil/pkg/prometheus/result/generic"
	"time"
)

type Result struct {
	Time  time.Time
	Value float64
	Raw   *generic.Result
}
