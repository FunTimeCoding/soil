package service_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/option"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/face"
)

func (o *Tester) daemonOption() *option.Memory {
	result := option.New()
	result.LitePath = o.Path
	result.HiddenTag = o.hiddenTag
	result.ReferenceRoot = o.referenceRoot
	result.Clock = o.clock
	result.Query = func() face.Client { return o.Indexer }

	return result
}
