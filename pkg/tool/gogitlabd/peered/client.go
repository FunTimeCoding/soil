package peered

import "github.com/funtimecoding/soil/pkg/gitlab/face"

func (p *Peered) Client() face.GitLabClient {
	return p.client
}
