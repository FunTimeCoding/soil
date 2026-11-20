package grant_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/client/mock_client"
	"testing"
)

type Tester struct {
	t          *testing.T
	Mock       *mock_client.Client
	Session    string
	Repository string
	Other      string
}
