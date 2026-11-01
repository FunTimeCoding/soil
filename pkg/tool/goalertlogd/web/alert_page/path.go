package alert_page

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/constant"
)

func Path(name string) string {
	return fmt.Sprintf(
		"%s?%s=%s",
		constant.AlertsPath,
		constant.NameParameter,
		name,
	)
}
