package example

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/argument"
	"github.com/funtimecoding/soil/pkg/console/status/option"
	"github.com/funtimecoding/soil/pkg/tool/common"
)

func AddResponder() {
	a := argument.NewSimple("add-responder")
	a.ParseSimple()
	r := a.RequiredPositional(0, "RESPONDER_NAME")
	c := common.Opsgenie()
	f := option.ExtendedColor.Copy()

	for _, a := range c.Open() {
		fmt.Println(a.Format(f))

		if false {
			c.AddResponderUser(a, r)
		}

		if true {
			break
		}
	}
}
