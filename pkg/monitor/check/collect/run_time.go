package collect

import (
	"github.com/funtimecoding/soil/pkg/bubbletea/model/monitor/fetch"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/monitor/collector"
	"github.com/funtimecoding/soil/pkg/time/constant"
	"k8s.io/apimachinery/pkg/util/duration"
	"time"
)

func runTime(
	s *collector.Collector,
	t time.Time,
) {
	console.Format(
		"%s %s %s\n",
		t.Format(constant.DateSecond),
		s.Name,
		duration.HumanDuration(s.Interval),
	)
	fetch.Run(s.Name)
}
