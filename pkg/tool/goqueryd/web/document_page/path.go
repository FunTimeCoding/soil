package document_page

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"strings"
)

func Path(virtualPath string) string {
	return fmt.Sprintf(
		"%s/%s",
		constant.DocumentsPath,
		strings.TrimPrefix(virtualPath, "qmd://"),
	)
}
