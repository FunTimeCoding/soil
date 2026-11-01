package session_page

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"net/url"
	"strconv"
)

func ListPath(
	active bool,
	page int,
) string {
	parameters := url.Values{}

	if active {
		parameters.Set(constant.Active, "true")
	}

	if page > 1 {
		parameters.Set(constant.Page, strconv.Itoa(page))
	}

	if len(parameters) == 0 {
		return constant.SessionsPath
	}

	return fmt.Sprintf("%s?%s", constant.SessionsPath, parameters.Encode())
}
