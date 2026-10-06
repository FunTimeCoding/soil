package response

import "github.com/funtimecoding/soil/pkg/provision/salt/basic/response/local_return"

type Local struct {
	Return []map[string]local_return.LocalReturn `json:"return"`
}
