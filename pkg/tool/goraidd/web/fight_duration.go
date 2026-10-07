package web

import (
	"github.com/funtimecoding/soil/pkg/raid/model/fight"
	"k8s.io/apimachinery/pkg/util/duration"
	"time"
)

func fightDuration(f fight.Fight) string {
	if f.DurationMS > 0 {
		return duration.HumanDuration(
			time.Duration(f.DurationMS) * time.Millisecond,
		)
	}

	return "-"
}
