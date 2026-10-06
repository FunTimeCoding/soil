package requester

import "github.com/funtimecoding/soil/pkg/web/requester/face"

func (r *Requester) WithAuthorizer(a face.Authorizer) *Requester {
	r.authorizer = a

	return r
}
