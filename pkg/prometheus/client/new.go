package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/prometheus/client_golang/api"
	"github.com/prometheus/client_golang/api/prometheus/v1"
	"net/http"
)

func New(
	locator string,
	r http.RoundTripper,
) v1.API {
	result, e := api.NewClient(api.Config{Address: locator, RoundTripper: r})
	errors.PanicOnError(e)

	return v1.NewAPI(result)
}
