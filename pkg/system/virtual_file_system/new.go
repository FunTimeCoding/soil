package virtual_file_system

import "github.com/funtimecoding/soil/pkg/system/virtual_file_system/file"

func New(options ...func(*System)) *System {
	result := &System{
		files:   make(map[string]*file.File),
		written: make(map[string]bool),
		deleted: make(map[string]bool),
	}

	for _, o := range options {
		o(result)
	}

	return result
}
