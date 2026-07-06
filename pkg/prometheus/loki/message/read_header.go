package message

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/web/telemetry/constant"
	"net/http"
	"strings"
)

func (m *Message) ReadHeader() http.Header {
	result := http.Header{}
	prefix := fmt.Sprintf("%s.", constant.HeaderPrefix)

	for k, v := range m.Values {
		if !strings.HasPrefix(k, prefix) {
			continue
		}

		key := strings.TrimPrefix(k, prefix)

		if slice, okayAny := v.([]any); okayAny && len(slice) > 0 {
			if s, okayString := slice[0].(string); okayString {
				result.Set(key, s)
			} else {
				result.Set(
					key,
					fmt.Sprintf("string fail: %v", slice[0]),
				)
			}
		} else {
			result.Set(key, fmt.Sprintf("any slice fail: %v", v))
		}
	}

	return result
}
