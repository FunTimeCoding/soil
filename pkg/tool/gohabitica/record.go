package gohabitica

import (
	"github.com/funtimecoding/soil/pkg/telemetry/constant"
	"github.com/funtimecoding/soil/pkg/telemetry/record"
)

func (x *Context) record(tool string) {
	x.Telemetry.Record(
		record.NewDomain(
			tool,
			constant.CommandLine,
			constant.User,
			constant.Success,
		),
	)
}
