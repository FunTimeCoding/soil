package face

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gosecretd/generated/client"
)

type Client interface {
	GetStatus(x context.Context) (*client.Status, error)
	ListSecrets(
		x context.Context,
		p *client.ListSecretsParams,
	) (*client.Listing, error)
	PutSecret(
		x context.Context,
		p *client.PutSecretParams,
		b client.PutSecretJSONRequestBody,
	) (*client.SecretKeys, error)
	RevealSecret(
		x context.Context,
		p *client.RevealSecretParams,
	) (*client.SecretValue, error)
	ShowSecret(
		x context.Context,
		p *client.ShowSecretParams,
	) (*client.SecretKeys, error)
	Reachable(x context.Context) error
	Close()
}
