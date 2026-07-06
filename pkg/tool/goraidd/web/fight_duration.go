package web

import (
	"github.com/funtimecoding/soil/pkg/raid"
	"k8s.io/apimachinery/pkg/util/duration"
	"time"
)

func fightDuration(f raid.Fight) string {
	if f.DurationMS > 0 {
		return duration.HumanDuration(
			time.Duration(f.DurationMS) * time.Millisecond,
		)
	}

	return "-"
}
