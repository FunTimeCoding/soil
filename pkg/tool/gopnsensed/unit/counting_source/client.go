package counting_source

import "github.com/funtimecoding/soil/pkg/opnsense/face"

type Client struct {
	face.OpnsenseClient
	reconfigured int
}
