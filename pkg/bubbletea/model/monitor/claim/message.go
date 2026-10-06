package claim

import "github.com/funtimecoding/soil/pkg/tool/gomonitord/generated/client"

type Message struct {
	Claims []client.Claim
	Error  error
}
