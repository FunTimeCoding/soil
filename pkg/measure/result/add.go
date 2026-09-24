package result

import "github.com/funtimecoding/soil/pkg/measure/file"

func (r *Result) Add(f *file.File) {
	r.Files = append(r.Files, f)
}
