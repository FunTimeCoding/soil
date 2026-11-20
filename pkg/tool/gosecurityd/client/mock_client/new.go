package mock_client

import (
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/generated/client"
	"testing"
	"time"
)

func New(t *testing.T) *Client {
	t.Helper()
	result := &Client{
		vulnerabilities: map[string]*client.Vulnerability{},
		statuses:        map[string]*client.DebianStatus{},
		snapshots:       map[string]client.SystemSnapshot{},
		updated:         map[string]time.Time{},
		details:         map[string]*client.SystemDetail{},
	}
	t.Cleanup(result.Close)

	return result
}
