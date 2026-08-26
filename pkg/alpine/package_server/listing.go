package package_server

import "github.com/funtimecoding/soil/pkg/alpine/index"

type Listing struct {
	Version      string         `json:"version"`
	Repository   string         `json:"repository"`
	Architecture string         `json:"architecture"`
	Packages     []*index.Entry `json:"packages"`
}
