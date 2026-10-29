package service_tester

import (
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/face"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/option"
)

func NewOption(
	path string,
	a face.Client,
) *option.Log {
	result := option.New()
	result.LitePath = path
	result.Alertmanager = func() face.Client { return a }

	return result
}
