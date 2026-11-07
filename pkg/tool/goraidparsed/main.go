package goraidparsed

import (
	"github.com/funtimecoding/soil/pkg/argument"
	"github.com/funtimecoding/soil/pkg/instrument"
	"github.com/funtimecoding/soil/pkg/tool/goraidparsed/constant"
	"github.com/funtimecoding/soil/pkg/tool/goraidparsed/option"
)

func Main() {
	s := instrument.New(constant.Identity)
	defer func() { s.Flush(recover()) }()
	a := argument.NewInstance(constant.Identity)
	a.Web()
	a.Parse()
	o := option.New()
	o.ParserPath = "/opt/parser"
	o.TemplatePath = "/opt/template/TW5_Top_Stat_Parse.html"
	o.OutputPath = "/srv/gw2-report"
	o.Read(a)
	Run(o, s)
}
