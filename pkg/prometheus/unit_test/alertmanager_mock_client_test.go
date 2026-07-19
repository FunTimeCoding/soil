package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/mock_client"
	"testing"
)

func TestAlertmanagerMockClient(t *testing.T) {
	assert.NotNil(t, mock_client.New())
}
