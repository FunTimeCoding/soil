package mark

import (
	"github.com/funtimecoding/soil/pkg/argument"
	argumentConstant "github.com/funtimecoding/soil/pkg/argument/constant"
	"github.com/funtimecoding/soil/pkg/generative/model_context/example/mark/option"
)

func Main() {
	a := argument.NewSimple("mark-example")
	a.Boolean(argumentConstant.Local, false, "Local STDIO")
	a.ParseSimple()
	o := option.New()
	o.Local = a.GetBoolean(argumentConstant.Local)
	Run(o)
}
