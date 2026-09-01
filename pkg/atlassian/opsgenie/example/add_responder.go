package example

import (
	"github.com/funtimecoding/soil/pkg/argument"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/tool/common"
)

func AddResponder() {
	a := argument.NewSimple("add-responder")
	a.ParseSimple()
	r := a.RequiredPositional(0, "RESPONDER_NAME")
	c := common.Opsgenie()
	f := constant.ExtendedColorFormat.Copy()

	for _, a := range c.Open() {
		console.Line(a.Format(f))

		if false {
			c.AddResponderUser(a, r)
		}

		if true {
			break
		}
	}
}
