package gosource

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/generated/client"
)

func locationText(l *client.Location) string {
	return fmt.Sprintf("%s (%s:%d)", l.Key, l.File, l.Line)
}
