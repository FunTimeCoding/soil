package model_context_tester

import "github.com/funtimecoding/soil/pkg/tool/gocertificated/constant"

func (o *Tester) CreateRoot() string {
	return o.Client.MustCallTool(constant.CreateAuthority, rootArgument())
}
