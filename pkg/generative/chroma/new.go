package chroma

import (
	"context"
	"github.com/amikos-tech/chroma-go/pkg/api/v2"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/web/locator"
)

func New(
	host string,
	port int,
) *Client {
	var o []v2.ClientOption

	if port == 0 {
		port = 8000
	}

	if host != "" {
		o = append(
			o,
			v2.WithBaseURL(
				locator.New(
					host,
				).Base("/api/v2").Port(port).Insecure().String(),
			),
		)
	}

	client, e := v2.NewHTTPClient(o...)
	errors.PanicOnError(e)

	return &Client{context: context.Background(), client: client}
}
