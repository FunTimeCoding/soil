package model_context

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/reacher/edge"
)

func recovered(e *edge.Edge) error {
	return fmt.Errorf("%s", e)
}
