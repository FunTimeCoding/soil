package round_tripper

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/prometheus/common/config"
	"net/http"
)

func New(
	user string,
	password string,
) http.RoundTripper {
	if user == "" && password == "" {
		return http.DefaultTransport
	}

	result, e := config.NewRoundTripperFromConfig(
		config.HTTPClientConfig{
			BasicAuth: &config.BasicAuth{
				Username: user,
				Password: config.Secret(password),
			},
		},
		"",
	)
	errors.PanicOnError(e)

	return result
}
