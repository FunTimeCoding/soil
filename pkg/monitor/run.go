package monitor

import (
	"fmt"
	argumentConstant "github.com/funtimecoding/soil/pkg/argument/constant"
	"github.com/funtimecoding/soil/pkg/constant"
	monitor "github.com/funtimecoding/soil/pkg/monitor/constant"
	"github.com/funtimecoding/soil/pkg/monitor/report"
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/system/run"
	"log"
	"time"
)

func Run(name string) *report.Report {
	r := run.New()
	r.Panic = false
	r.Start(name, fmt.Sprintf("--%s", argumentConstant.Notation))
	result := report.New()

	if r.Error != nil {
		s := fmt.Sprintf("run fail: %s %s", name, r.Error)
		log.Print(s)
		result.AddItem(
			monitor.MonitorCollector,
			monitor.MonitorCollector.StringIdentifier(name),
			constant.Critical,
			s,
			"",
			&time.Time{},
		)

		return result
	}

	if e := notation.Decode(r.OutputString, &result); e != nil {
		log.Printf("parse fail: %s %s", name, e)
		result.AddItem(
			monitor.MonitorCollector,
			monitor.MonitorCollector.StringIdentifier(name),
			constant.Critical,
			fmt.Sprintf("parse fail: %s %s", name, r.Error),
			"",
			&time.Time{},
		)

		return result
	}

	return result
}
