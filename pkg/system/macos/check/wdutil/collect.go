package wdutil

import (
	"github.com/funtimecoding/soil/pkg/system/constant"
	"github.com/funtimecoding/soil/pkg/system/run"
)

func collect() *Result {
	r := run.New()
	r.Start("sudo", constant.Wdutil, constant.WdutilInformation)

	return &Result{Sequence: parseKey(r.OutputString, "Channel Sequence")}
}
