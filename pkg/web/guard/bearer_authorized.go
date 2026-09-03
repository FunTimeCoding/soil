package guard

import (
	"crypto/subtle"
	"github.com/funtimecoding/soil/pkg/strings/join/key_value"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"net/http"
)

func bearerAuthorized(
	q *http.Request,
	tokens []string,
) bool {
	h := []byte(q.Header.Get(constant.Authorization))

	for _, token := range tokens {
		expected := []byte(key_value.Space(constant.Bearer, token))

		if subtle.ConstantTimeCompare(h, expected) == 1 {
			return true
		}
	}

	return false
}
