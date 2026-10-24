package requester

import (
	"github.com/funtimecoding/soil/pkg/web/locator"
	"github.com/funtimecoding/soil/pkg/web/requester/face"
	"golang.org/x/time/rate"
	"net/http"
	"time"
)

type Requester struct {
	base       *locator.Locator
	client     *http.Client
	userAgent  string
	header     map[string]string
	authorizer face.Authorizer
	refusal    func(int, []byte) error
	backoff    time.Duration
	pace       *rate.Limiter
}
