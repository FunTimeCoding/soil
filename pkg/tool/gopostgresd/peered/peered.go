package peered

import (
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/postgres/face"
)

type Peered struct {
	client face.PostgresClient
	guard  *peer.Guard
}
