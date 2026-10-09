package worker

import (
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/types/floor"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/types/floor/guest"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/types/floor/node"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/types/floor/storage"
	"slices"
	"strings"
	"time"
)

func (w *Worker) Poll() {
	f := floor.New()

	for _, i := range w.service.Instances() {
		start := time.Now()
		e := w.pollInstance(i.Name, f)

		if e != nil {
			w.logFailure(i.Name, e)
			w.collector.Clear(i.Name)
			w.collector.SetScrape(i.Name, false, time.Since(start))

			continue
		}

		if edge := w.reacher.Observe(i.Name, nil); edge != nil {
			w.log.Plain("poll hypervisor %s", edge)
		}

		w.collector.SetScrape(i.Name, true, time.Since(start))
	}

	slices.SortFunc(
		f.Nodes,
		func(
			a node.Node,
			b node.Node,
		) int {
			if c := strings.Compare(a.Hypervisor, b.Hypervisor); c != 0 {
				return c
			}

			return strings.Compare(a.Name, b.Name)
		},
	)
	slices.SortFunc(
		f.Guests,
		func(
			a guest.Guest,
			b guest.Guest,
		) int {
			if c := strings.Compare(a.Hypervisor, b.Hypervisor); c != 0 {
				return c
			}

			if c := strings.Compare(a.Node, b.Node); c != 0 {
				return c
			}

			return strings.Compare(a.Name, b.Name)
		},
	)
	slices.SortFunc(
		f.Storages,
		func(
			a storage.Storage,
			b storage.Storage,
		) int {
			if c := strings.Compare(a.Hypervisor, b.Hypervisor); c != 0 {
				return c
			}

			return strings.Compare(a.Name, b.Name)
		},
	)
	w.mutex.Lock()
	previous := w.floor
	w.floor = f
	w.mutex.Unlock()

	if previous == nil || !previous.Equal(*f) {
		w.notifier.Notify()
	}
}
