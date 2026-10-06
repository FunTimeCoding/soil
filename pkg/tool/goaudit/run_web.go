package goaudit

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goaudit/format"
	"github.com/funtimecoding/soil/pkg/tool/goaudit/scan/matrix"
)

func runWeb(frontends []*matrix.Frontend) {
	fmt.Print(format.Frontends(frontends))
}
