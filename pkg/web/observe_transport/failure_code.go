package observe_transport

import (
	"context"
	"errors"
	"github.com/funtimecoding/soil/pkg/errors/connection"
	"github.com/funtimecoding/soil/pkg/web/constant"
)

func failureCode(e error) string {
	if errors.Is(e, context.Canceled) {
		return constant.MetricCanceledCode
	}

	if f := connection.Classify(e); f != nil {
		return f.Kind
	}

	return constant.MetricErrorCode
}
