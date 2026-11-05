package browser

import "github.com/funtimecoding/soil/pkg/chromium/face"

func New(c face.ChromiumClient) *Browser {
	return &Browser{ChromiumClient: c}
}
