package decision_page

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/constant"
	"net/url"
	"strconv"
)

func actionPath(
	path string,
	identifier uint,
	kind constant.AnswerKind,
) string {
	parameters := url.Values{}
	parameters.Set(
		constant.IdentifierParameter,
		strconv.FormatUint(uint64(identifier), 10),
	)

	if kind != "" {
		parameters.Set(constant.KindParameter, string(kind))
	}

	return fmt.Sprintf("%s?%s", path, parameters.Encode())
}
