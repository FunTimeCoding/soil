package browser

import "github.com/funtimecoding/soil/pkg/browser/face"

func New(c face.ChromiumClient) *Browser {
	return &Browser{ChromiumClient: c}
}
