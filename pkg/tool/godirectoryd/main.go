package godirectoryd

import (
	"github.com/funtimecoding/soil/pkg/argument"
	directoryConstant "github.com/funtimecoding/soil/pkg/directory/constant"
	"github.com/funtimecoding/soil/pkg/instrument"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/option"
)

func Main() {
	i := instrument.New(constant.Identity)
	defer func() { i.Flush(recover()) }()
	a := argument.NewInstance(constant.Identity)
	a.Web()
	a.Parse()
	o := option.New()
	o.Read(a)
	o.Host = environment.Required(directoryConstant.HostEnvironment)
	Run(o, i)
}
