package graph_tester

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/microsoft/constant"
)

func TokenPath() string {
	return fmt.Sprintf(constant.TokenPath, "tenant")
}
