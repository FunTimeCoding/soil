package goalpined

import (
	"github.com/funtimecoding/soil/pkg/argument"
	"github.com/funtimecoding/soil/pkg/instrument"
	"github.com/funtimecoding/soil/pkg/tool/goalpined/constant"
	"github.com/funtimecoding/soil/pkg/tool/goalpined/option"
)

func Main(
	version string,
	gitHash string,
	buildDate string,
) {
	i := instrument.New(constant.Identity, version)
	defer func() { i.Flush(recover()) }()
	a := argument.NewInstance(constant.Identity)
	a.Web()
	a.Parse(version, gitHash, buildDate)
	o := option.New()
	o.Address = a.Address()
	o.Version = version
	Run(o, i)
}
