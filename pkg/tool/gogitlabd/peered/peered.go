package peered

import (
	"github.com/funtimecoding/soil/pkg/gitlab/face"
	"github.com/funtimecoding/soil/pkg/peer"
)

type Peered struct {
	client face.GitLabClient
	guard  *peer.Guard
}
