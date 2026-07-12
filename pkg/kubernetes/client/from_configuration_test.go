package client

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"k8s.io/client-go/rest"
	"testing"
)

func TestFromConfiguration(t *testing.T) {
	c, e := fromConfiguration(
		&rest.Config{Host: "https://127.0.0.1:1"},
		"test",
	)
	assert.FatalOnError(t, e)
	assert.NotNil(t, c.context)
	assert.NotNil(t, c.configuration)
	assert.NotNil(t, c.client)
	assert.NotNil(t, c.metric)
	assert.NotNil(t, c.dynamic)
	assert.NotNil(t, c.clients)
	assert.String(t, "test", c.cluster)
}
