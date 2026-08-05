package server

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/constant"
)

func snippetPath(name string) string {
	return fmt.Sprintf("%s/%s", constant.SnippetDirectory, name)
}
