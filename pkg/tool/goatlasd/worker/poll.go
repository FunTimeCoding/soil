package worker

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/gazetteer"
)

func (w *Worker) poll() {
	x := context.Background()
	s, e := gazetteer.Load(x, w.netbox)
	errors.PanicOnError(e)
	w.Collect(x, s)
}
