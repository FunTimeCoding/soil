package watcher

import "github.com/funtimecoding/soil/pkg/tool/gosproutd/service"

func (w *Watcher) scan() {
	walked := w.walkDirectory()
	var files []service.DiscoveredFile

	for _, f := range walked {
		files = append(
			files,
			service.DiscoveredFile{
				Name:        f.name,
				Path:        f.path,
				ContentHash: f.contentHash,
				Content:     f.content,
			},
		)
	}

	w.service.Sync(files)
}
