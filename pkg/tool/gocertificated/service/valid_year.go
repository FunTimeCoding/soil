package service

import "github.com/funtimecoding/soil/pkg/tool/gocertificated/constant"

func validYear(
	given *int,
	kind constant.CertificateKind,
) int {
	if given != nil && *given > 0 {
		return *given
	}

	if kind == constant.KindRoot {
		return constant.RootValidityYear
	}

	return constant.IntermediateValidityYear
}
