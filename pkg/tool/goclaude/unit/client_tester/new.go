package client_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/client"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/face"
	"testing"
)

func New(
	t *testing.T,
	address string,
) face.Client {
	t.Helper()
	result := client.New(address, false, "")
	t.Cleanup(result.Close)

	return result
}
