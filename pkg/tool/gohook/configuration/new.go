package configuration

import "github.com/funtimecoding/soil/pkg/tool/gohook/job"

func New(
	root string,
	path string,
) *Configuration {
	return &Configuration{
		Root:  root,
		Path:  path,
		Hooks: map[string][]*job.Job{},
	}
}
