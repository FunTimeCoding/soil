package worker_tester

import "github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"

func (o *Sink) Add(v client.NotifyRequest) {
	o.mutex.Lock()
	defer o.mutex.Unlock()
	o.received = append(o.received, v)
}
