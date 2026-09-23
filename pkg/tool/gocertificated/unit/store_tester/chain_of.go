package store_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/authority"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/store"
)

func ChainOf(s *store.Store) (*authority.Authority, *authority.Authority) {
	return authority.New(s.MustAuthority(constant.RootAuthority).Material()),
		authority.New(
			s.MustAuthority(constant.FixtureClusterAuthority).Material(),
		)
}
