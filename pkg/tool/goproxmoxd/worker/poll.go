package worker

import (
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/types/floor"
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
			w.log.Plain("poll hypervisor %s failed: %v", i.Name, e)
			w.collector.Clear(i.Name)
			w.collector.SetScrape(i.Name, false, time.Since(start))

			continue
		}

		w.collector.SetScrape(i.Name, true, time.Since(start))
	}

	slices.SortFunc(
		f.Nodes,
		func(a floor.Node, b floor.Node) int {
			if c := strings.Compare(a.Hypervisor, b.Hypervisor); c != 0 {
				return c
			}

			return strings.Compare(a.Name, b.Name)
		},
	)
	slices.SortFunc(
		f.Guests,
		func(a floor.Guest, b floor.Guest) int {
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
		func(a floor.Storage, b floor.Storage) int {
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
