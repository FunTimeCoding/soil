package proxmox

import (
	"context"
	"github.com/funtimecoding/soil/pkg/proxmox/constant"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"github.com/funtimecoding/soil/pkg/web/stall_client"
	"github.com/luthermonson/go-proxmox"
)

// Reference: https://pve.proxmox.com/pve-docs/api-viewer
func New(
	host string,
	o ...Option,
) *Client {
	result := &Client{context: context.Background()}

	for _, p := range o {
		p(result)
	}

	var option []proxmox.Option

	if result.user != "" && result.password != "" {
		option = append(
			option,
			proxmox.WithCredentials(
				&proxmox.Credentials{
					Username: result.user,
					Password: result.password,
				},
			),
		)
	} else if result.token != "" && result.secret != "" {
		option = append(
			option,
			proxmox.WithAPIToken(result.token, result.secret),
		)
	} else {
		panic("missing credentials")
	}

	if result.log {
		option = append(
			option,
			proxmox.WithLogger(
				&proxmox.LeveledLogger{Level: proxmox.LevelDebug},
			),
		)
	}

	b := stall_client.New().Timeout(result.timeout)

	if result.untrusted {
		b.Insecure()
	}

	if result.verbose {
		b.Verbose()
	}

	result.pool = b.Build()
	option = append(option, proxmox.WithHTTPClient(result.pool.Client()))
	p := constant.Port

	if result.port > 0 {
		p = result.port
	}

	result.client = proxmox.NewClient(
		locator.New(host).Port(p).Base(constant.Base).String(),
		option...,
	)

	return result
}
