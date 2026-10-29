package option

import (
	"github.com/funtimecoding/soil/pkg/chromium"
	"github.com/funtimecoding/soil/pkg/tool/gochromed/browser"
	"github.com/funtimecoding/soil/pkg/tool/gochromed/face"
)

func New() *Chrome {
	return &Chrome{
		Source: func() face.ChromiumSource {
			return browser.New(chromium.NewEnvironment())
		},
	}
}
