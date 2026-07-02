package example

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/ssh/tunnel"
	"github.com/funtimecoding/go-library/pkg/ssh/tunnel/constant"
	"github.com/funtimecoding/go-library/pkg/system"
)

func Forward() {
	a := argument.NewSimple("tunnel")
	a.String(argument.Host, "", "Relay host")
	a.String(constant.TargetHost, "", "Target host")
	a.Integer(constant.TargetPort, 0, "Target port")
	a.ParseSimple()
	t := tunnel.New()

	if false {
		t.Verbose = true
		t.NoOutput = true
	}

	fmt.Printf("Start: %+v\n", t)
	t.Start(
		a.Required(argument.Host),
		a.Required(constant.TargetHost),
		a.RequiredInteger(constant.TargetPort),
		0,
	)
	defer t.Stop()
	fmt.Println("Sleep 10")
	system.Sleep(10)
	fmt.Println("Stop")
}
