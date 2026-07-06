package time_string

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
	"time"
)

func TestOnlyLatest(t *testing.T) {
	assert.Count(
		t,
		1,
		OnlyLatest(
			[]*Result{
				{
					Metric: "",
					Time:   time.Now(),
					Value:  "",
				},
			},
		),
	)
}
