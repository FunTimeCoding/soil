package example

import (
	"github.com/funtimecoding/soil/pkg/argument"
	argumentConstant "github.com/funtimecoding/soil/pkg/argument/constant"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/ssh/constant"
	"github.com/funtimecoding/soil/pkg/ssh/tunnel"
	"github.com/funtimecoding/soil/pkg/system"
)

func Forward() {
	a := argument.NewSimple("tunnel")
	a.String(argumentConstant.Host, "", "Relay host")
	a.String(constant.TargetHost, "", "Target host")
	a.Integer(constant.TargetPort, 0, "Target port")
	a.ParseSimple()
	t := tunnel.New()

	if false {
		t.Verbose = true
		t.NoOutput = true
	}

	console.Format("Start: %+v\n", t)
	t.Start(
		a.Required(argumentConstant.Host),
		a.Required(constant.TargetHost),
		a.RequiredInteger(constant.TargetPort),
		0,
	)
	defer t.Stop()
	console.Line("Sleep 10")
	system.Sleep(10)
	console.Line("Stop")
}
