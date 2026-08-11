package goaudit

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goaudit/format"
	"github.com/funtimecoding/soil/pkg/tool/goaudit/scan"
)

func runWeb(frontends []*scan.Frontend) {
	fmt.Print(format.Frontends(frontends))
}
