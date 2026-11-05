package goprocess

import (
	"github.com/funtimecoding/soil/pkg/argument"
	"github.com/funtimecoding/soil/pkg/instrument"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/goprocess/constant"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/client"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/face"
)

func Main() {
	s := instrument.NewCommandLine(constant.Identity)
	defer func() { s.Flush(recover()) }()
	o := New(
		func(procfile string) face.Client { return client.NewLocal(procfile) },
		terminal.New(s),
	)
	argument.CobraInstrument(o, s)
	argument.CobraStamp(o, constant.Identity)
	argument.CobraExecute(o)
}
