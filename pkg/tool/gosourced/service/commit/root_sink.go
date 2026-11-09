package commit

import "github.com/funtimecoding/soil/pkg/tool/gosourced/sink"

func RootSink(
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
