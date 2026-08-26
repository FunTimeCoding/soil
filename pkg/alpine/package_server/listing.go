package package_server

import "github.com/funtimecoding/soil/pkg/alpine/index"

type Listing struct {
	Version      string
	Repository   string
	Architecture string
	Packages     []*index.Entry
}
