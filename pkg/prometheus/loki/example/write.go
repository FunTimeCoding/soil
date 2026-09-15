package example

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/prometheus/loki"
	"github.com/funtimecoding/soil/pkg/time/constant"
	"time"
)

func Write() {
	c := loki.NewEnvironment(true)
	c.Push(map[string]string{"application": "example"}, "test message")
	end := time.Now()
	r, _ := c.QueryRange(
		`{application="example"}`,
		end.Add(-time.Hour),
		end,
		10,
	)

	for _, v := range r {
		console.Format("%s %s\n", v.Time.Format(constant.DateMinute), v.Text)
	}
}
