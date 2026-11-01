package session_page

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/constant"
	"net/url"
)

func Path(name string) string {
	parameters := url.Values{}
	parameters.Set(constant.SessionParameter, name)

	return fmt.Sprintf("%s?%s", constant.SessionPath, parameters.Encode())
}
