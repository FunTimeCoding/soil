package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/kubernetes/client"
	"k8s.io/client-go/rest"
	"testing"
)

func TestFromConfiguration(t *testing.T) {
	c, e := client.FromConfiguration(
		&rest.Config{Host: "https://127.0.0.1:1"},
		"test",
	)
	assert.FatalOnError(t, e)
	assert.Nil(t, c.Validate())
	assert.String(t, "test", c.Cluster())
}
