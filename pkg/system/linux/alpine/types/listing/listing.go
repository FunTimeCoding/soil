package listing

import "github.com/funtimecoding/soil/pkg/system/linux/alpine/index"

type Listing struct {
	Version      string
	Repository   string
	Architecture string
	Packages     []*index.Entry
}
