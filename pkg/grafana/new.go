package grafana

import (
	"fmt"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/go-openapi/strfmt"
	"github.com/grafana/grafana-openapi-client-go/client"
)

func New(
	host string,
	port string,
	token string,
) *Client {
	return &Client{
		client: client.NewHTTPClientWithConfig(
			strfmt.Default,
			&client.TransportConfig{
				Host:     fmt.Sprintf("%s:%s", host, port),
				BasePath: "/api",
				Schemes:  []string{webConstant.Secure},
				APIKey:   token,
				OrgID:    1,
			},
		),
	}
}
