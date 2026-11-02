package collect

import (
	"github.com/funtimecoding/soil/pkg/console"
	timeConstant "github.com/funtimecoding/soil/pkg/time/constant"
	"time"
)

func loopCheck(t time.Time) {
	console.Format("Time: %s\n", t.Format(timeConstant.DateMinute))
	Check(false, false)
}
