package notifier

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/connector"
)

type Notifier struct {
	connector *connector.Client
	source    string
	reporter  face.Reporter
}
