package notifier

import (
	"github.com/funtimecoding/soil/pkg/face"
	goclauded "github.com/funtimecoding/soil/pkg/tool/goclauded/face"
)

type Notifier struct {
	connector goclauded.Client
	source    string
	reporter  face.Reporter
}
