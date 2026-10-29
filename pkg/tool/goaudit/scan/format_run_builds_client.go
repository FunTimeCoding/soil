package scan

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goaudit/constant"
)

func formatRunBuildsClient(name string) string {
	return fmt.Sprintf(constant.RunBuildsClientText, name)
}
