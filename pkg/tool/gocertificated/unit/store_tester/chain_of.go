package store_tester

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/authority"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/service"
)

func ChainOf(v *service.Service) (*authority.Authority, *authority.Authority) {
	root, e := v.Authority(constant.RootAuthority)
	errors.PanicOnError(e)
	cluster, f := v.Authority(constant.FixtureClusterAuthority)
	errors.PanicOnError(f)

	return authority.New(root.Material()), authority.New(cluster.Material())
}
