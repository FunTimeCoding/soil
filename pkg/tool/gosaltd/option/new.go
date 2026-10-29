package option

import "github.com/funtimecoding/soil/pkg/provision/salt"

func New() *Salt {
	return &Salt{Master: salt.NewEnvironment}
}
