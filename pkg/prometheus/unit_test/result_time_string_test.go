package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/prometheus/result/time_string"
	"testing"
	"time"
)

func TestOnlyLatest(t *testing.T) {
	assert.Count(
		t,
		1,
		time_string.OnlyLatest(
			[]*time_string.Result{
				{
					Metric: "",
					Time:   time.Now(),
					Value:  "",
				},
			},
		),
	)
}
