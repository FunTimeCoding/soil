package collect

import (
	"fmt"
	library "github.com/funtimecoding/soil/pkg/time"
	"time"
)

func loopCheck(t time.Time) {
	fmt.Printf(
		"Time: %s\n",
		t.Format(library.DateMinute),
	)
	Check(false, false)
}
