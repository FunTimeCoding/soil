package collect

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/time/constant"
	"time"
)

func loopCheck(t time.Time) {
	fmt.Printf("Time: %s\n", t.Format(constant.DateMinute))
	Check(false, false)
}
