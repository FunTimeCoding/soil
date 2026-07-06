package project

import "github.com/funtimecoding/soil/pkg/console/description"

func (p *Project) Meta() *description.Description {
	return description.New("Project", "Project")
}
