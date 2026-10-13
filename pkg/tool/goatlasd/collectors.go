package goatlasd

import (
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/collector/kubernetes"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/collector/label"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/collector/outpost"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/collector/process"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/face"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"net/http"
)

func collectors(
	n *client.ClientWithResponses,
	l *logger.Logger,
	h *http.Client,
) []face.Collector {
	result := []face.Collector{kubernetes.NewEnvironment(), label.New(n)}

	if p := process.NewOptional(h); p != nil {
		result = append(result, p)
	}

	if o := outpost.NewOptional(l, h); o != nil {
		result = append(result, o)
	}

	return result
}
