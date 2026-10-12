package gosecret

import "github.com/funtimecoding/soil/pkg/tool/gosecretd/generated/client"

type interfaceStore struct {
	client *client.ClientWithResponses
}
