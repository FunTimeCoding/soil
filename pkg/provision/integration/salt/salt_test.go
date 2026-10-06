//go:build local

package salt

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/provision/salt"
	"testing"
)

func TestKeysMinionsAndJobsRead(t *testing.T) {
	c := salt.NewEnvironment()
	keys, e := c.Keys()
	assert.FatalOnError(t, e)
	assert.NotEmpty(t, keys.Minions)
	minions, f := c.Minions()
	assert.FatalOnError(t, f)
	assert.NotEmpty(t, minions)
	_, g := c.Jobs()
	assert.FatalOnError(t, g)
}
