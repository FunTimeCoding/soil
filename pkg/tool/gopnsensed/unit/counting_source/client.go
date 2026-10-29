package counting_source

import "github.com/funtimecoding/soil/pkg/tool/gopnsensed/face"

type Client struct {
	face.OpnsenseSource
	reconfigured int
}
