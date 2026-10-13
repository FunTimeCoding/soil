package ollama

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/generative/constant"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/stall_client"
	"github.com/ollama/ollama/api"
	"net/url"
)

// Reference: https://github.com/ollama/ollama/blob/main/docs/api.md
func New(o ...Option) *Client {
	result := &Client{context: context.Background()}

	for _, p := range o {
		p(result)
	}

	if result.host == "" {
		result.host = constant.OllamaHost
	}

	if result.port == 0 {
		result.port = constant.OllamaPort
	}

	var scheme string

	if result.secure {
		scheme = webConstant.Secure
	} else {
		scheme = webConstant.Insecure
	}

	result.pool = stall_client.New().Long().Build()
	result.client = api.NewClient(
		&url.URL{
			Scheme: scheme,
			Host:   fmt.Sprintf("%s:%d", result.host, result.port),
		},
		result.pool.Client(),
	)

	return result
}
