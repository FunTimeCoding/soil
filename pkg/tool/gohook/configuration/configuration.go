package configuration

import "github.com/funtimecoding/soil/pkg/tool/gohook/job"

type Configuration struct {
	Path  string
	Root  string
	Hooks map[string][]*job.Job
}
