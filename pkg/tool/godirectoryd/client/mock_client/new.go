package mock_client

import "github.com/funtimecoding/soil/pkg/failure"

func New() *Client {
	result := &Client{Base: "dc=example,dc=org", Passwords: map[string]string{}}
	result.failures = failure.New(result.GetUser)

	return result
}
