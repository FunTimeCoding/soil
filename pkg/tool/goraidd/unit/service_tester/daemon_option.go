package service_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/goraidd/option"
	"github.com/funtimecoding/soil/pkg/tool/goraidparsed/face"
)

func (o *Tester) daemonOption() *option.Raid {
	result := option.New()
	result.LitePath = o.litePath
	result.LogCachePath = o.LogCache
	result.ElitePath = o.Elite
	result.OutputPath = o.Output
	result.Clock = o.clock
	result.Parser = func() face.Client { return o.Parser }

	return result
}
