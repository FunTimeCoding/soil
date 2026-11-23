package gosalt

import (
	"github.com/funtimecoding/soil/pkg/argument"
	"github.com/funtimecoding/soil/pkg/instrument"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/gosalt/constant"
)

func Main() {
	s := instrument.NewCommandLine(constant.Identity)
	defer func() { s.Flush(recover()) }()
	o := New(terminal.New(s))
	argument.CobraInstrument(o, s)
	argument.CobraStamp(o, constant.Identity)
	argument.CobraExecute(o)
}
