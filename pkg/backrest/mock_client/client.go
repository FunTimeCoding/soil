package mock_client

import "github.com/funtimecoding/soil/pkg/backrest/response"

type Client struct {
	Failure error
	Calls   int
	stanzas []*response.Stanza
}
