package console_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
	"testing"
)

type Tester struct {
	t      *testing.T
	client *client.ClientWithResponses
}
