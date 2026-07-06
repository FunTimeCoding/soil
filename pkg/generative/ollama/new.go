package ollama

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/generative/ollama/constant"
	web "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/ollama/ollama/api"
	"net/http"
	"net/url"
)

func New(o ...Option) *Client {
	result := &Client{context: context.Background()}

	for _, p := range o {
		p(result)
	}

	if result.host == "" {
		result.host = constant.Host
	}

	if result.port == 0 {
		result.port = constant.Port
	}

	var scheme string

	if result.secure {
		scheme = web.Secure
	} else {
		scheme = web.Insecure
	}

	// https://github.com/ollama/ollama/blob/main/docs/api.md
	result.client = api.NewClient(
		&url.URL{
			Scheme: scheme,
			Host:   fmt.Sprintf("%s:%d", result.host, result.port),
		},
		http.DefaultClient,
	)

	return result
}
