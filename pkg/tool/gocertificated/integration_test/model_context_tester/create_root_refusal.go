package model_context_tester

import "github.com/funtimecoding/soil/pkg/tool/gocertificated/constant"

func (o *Tester) CreateRootRefusal() string {
	return o.Client.MustCallToolError(constant.CreateAuthority, rootArgument())
}
