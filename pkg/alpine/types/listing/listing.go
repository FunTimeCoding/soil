package listing

import "github.com/funtimecoding/soil/pkg/alpine/index"

type Listing struct {
	Version      string
	Repository   string
	Architecture string
	Packages     []*index.Entry
}
