package service

import "github.com/funtimecoding/soil/pkg/tool/gocertificated/constant"

func validDay(given *int) int {
	if given != nil && *given > 0 {
		return *given
	}

	return constant.LeafValidityDay
}
