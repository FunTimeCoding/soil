package pool

import "github.com/funtimecoding/soil/pkg/face"

func release(c any) {
	if v, okay := c.(face.Closer); okay {
		v.Close()
	}
}
