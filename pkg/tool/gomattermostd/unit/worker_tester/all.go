package worker_tester

import "github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"

func (o *Sink) All() []client.NotifyRequest {
	o.mutex.Lock()
	defer o.mutex.Unlock()

	return append([]client.NotifyRequest{}, o.received...)
}
