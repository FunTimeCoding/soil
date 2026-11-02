package service_tester

import (
	"github.com/funtimecoding/soil/pkg/prometheus/face"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/option"
)

func NewOption(
	path string,
	a face.AlertmanagerClient,
) *option.Log {
	result := option.New()
	result.LitePath = path
	result.Alertmanager = func() face.AlertmanagerClient { return a }

	return result
}
