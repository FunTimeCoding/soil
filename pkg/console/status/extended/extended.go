package extended

import "github.com/funtimecoding/soil/pkg/console/types/raw"

type Extended struct {
	Identifier  int
	Name        string
	Description string
	Raw         *raw.Raw
}
