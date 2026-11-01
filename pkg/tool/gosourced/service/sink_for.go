package service

import "github.com/funtimecoding/soil/pkg/tool/gosourced/service/sink"

func sinkFor(
	outs []*sink.Sink,
	root string,
) *sink.Sink {
	for _, o := range outs {
		if o.Root() == root {
			return o
		}
	}

	return nil
}
