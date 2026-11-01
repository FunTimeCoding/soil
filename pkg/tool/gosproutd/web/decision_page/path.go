package decision_page

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/constant"
	"net/url"
	"strconv"
)

func Path(
	identifier uint,
	open bool,
) string {
	parameters := url.Values{}
	parameters.Set(
		constant.IdentifierParameter,
		strconv.FormatUint(uint64(identifier), 10),
	)

	if open {
		parameters.Set(constant.OpenParameter, "1")
	}

	return fmt.Sprintf("%s?%s", constant.DecisionPath, parameters.Encode())
}
