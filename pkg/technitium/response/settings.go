package response

import "github.com/funtimecoding/soil/pkg/technitium/key"

type Settings struct {
	Keys []*key.Key `json:"tsigKeys"`
}
