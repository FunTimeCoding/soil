package graph_tester

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/microsoft/constant"
)

func DevicePath() string {
	return fmt.Sprintf(constant.DeviceCodePath, "tenant")
}
