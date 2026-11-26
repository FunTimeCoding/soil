package goladybirdd

import (
	"github.com/funtimecoding/soil/pkg/argument"
	"github.com/funtimecoding/soil/pkg/instrument"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/tool/goladybirdd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goladybirdd/option"
)

func Main() {
	s := instrument.New(constant.Identity)
	defer func() { s.Flush(recover()) }()
	a := argument.NewInstance(constant.Identity)
	a.Web()
	a.Boolean(
		constant.VisibleArgument,
		false,
		"Show the browser window instead of running headless",
	)
	a.Parse()
	o := option.New()
	o.Read(a)
	o.Headless = !a.GetBoolean(constant.VisibleArgument)
	o.DownloadDirectory = environment.Required(
		constant.DownloadDirectoryEnvironment,
	)
	Run(o, s)
}
