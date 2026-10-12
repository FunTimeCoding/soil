package gosecret

import "github.com/funtimecoding/soil/pkg/tool/gosecretd/generated/client"

func newInterfaceStore(c *client.ClientWithResponses) *interfaceStore {
	return &interfaceStore{client: c}
}
