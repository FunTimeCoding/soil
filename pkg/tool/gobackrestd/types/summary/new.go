package summary

import "github.com/funtimecoding/soil/pkg/backrest/response"

func New(
	stanza string,
	status int,
) *Summary {
	return &Summary{
		Stanza: stanza,
		Status: status,
		Latest: make(map[string]*response.Backup),
	}
}
