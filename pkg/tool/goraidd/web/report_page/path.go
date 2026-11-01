package report_page

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/constant"
)

func Path(name string) string {
	return fmt.Sprintf("%s/%s", constant.ReportsPath, name)
}
