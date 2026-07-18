package convert

import (
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/store/label"
)

func Label(v *label.Label) *server.Label {
	return &server.Label{Key: v.Key, Value: v.Value}
}
