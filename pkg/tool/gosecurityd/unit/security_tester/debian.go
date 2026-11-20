package security_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/service/debian"
)

func (o *Tester) Debian() *debian.Service {
	o.t.Helper()

	if o.debian != nil {
		return o.debian
	}

	o.debian = debian.New(
		o.Source(
			o.SecurityTracker,
			constant.DebianCache,
			true,
			constant.SupportedReleases...,
		),
		o.Source(
			o.ExtendedTracker,
			constant.ExtendedCache,
			false,
			constant.ExtendedReleases...,
		),
	)

	return o.debian
}
