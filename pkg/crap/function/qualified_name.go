package function

import "github.com/funtimecoding/soil/pkg/strings/join"

func (f *Function) QualifiedName() string {
	if f.Receiver == "" {
		return f.Name
	}

	return join.Dot([]string{f.Receiver, f.Name})
}
