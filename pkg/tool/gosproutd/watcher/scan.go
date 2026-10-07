package watcher

import "github.com/funtimecoding/soil/pkg/tool/gosproutd/types/discovered_file"

func (w *Watcher) scan() {
	walked := w.walkDirectory()
	var files []discovered_file.File

	for _, f := range walked {
		files = append(
			files,
			discovered_file.File{
				Name:        f.Name,
				Path:        f.Path,
				ContentHash: f.ContentHash,
				Content:     f.Content,
				ModifiedAt:  f.ModifiedAt,
			},
		)
	}

	w.service.Sync(files)
}
