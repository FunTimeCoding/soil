package service

import (
	"github.com/funtimecoding/soil/pkg/errors/timeout"
	"github.com/funtimecoding/soil/pkg/relational/lite"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
)

func classifyWrite(e error) error {
	if lite.IsBusy(e) {
		return timeout.Format(constant.StoreBusy, e)
	}

	return e
}
