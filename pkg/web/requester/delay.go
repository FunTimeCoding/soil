package requester

import (
	"github.com/funtimecoding/soil/pkg/web/constant"
	"net/http"
	"strconv"
	"time"
)

func (r *Requester) delay(
	attempt int,
	response *http.Response,
) time.Duration {
	if response != nil {
		seconds, e := strconv.Atoi(response.Header.Get(constant.RetryAfter))

		if e == nil && seconds >= 0 {
			return min(
				time.Duration(seconds)*time.Second,
				constant.MaximumRetryAfter,
			)
		}
	}

	return r.backoff * time.Duration(attempt+1)
}
