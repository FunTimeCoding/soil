package console

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/system"
	"net/http"
)

func Emit(
	content string,
	status int,
) {
	if status >= http.StatusBadRequest {
		system.Exitln(content)
	}

	if content == "" {
		return
	}

	fmt.Println(content)
}
