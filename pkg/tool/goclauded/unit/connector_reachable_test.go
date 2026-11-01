package unit

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/client"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/base"
	"testing"
)

func TestConnectorReachesTheServer(t *testing.T) {
	s := base.New(t)
	assert.FatalOnError(t, s.Connector(t).Reachable(context.Background()))
}

func TestConnectorReachableRefusesAWrongToken(t *testing.T) {
	s := base.New(t)
	c := client.New(fmt.Sprintf("http://localhost:%d", s.Port), false, "wrong")
	defer c.Close()
	assert.Error(t, c.Reachable(context.Background()))
}

func TestConnectorAddressNamesHostAndPort(t *testing.T) {
	s := base.New(t)
	assert.String(
		t,
		fmt.Sprintf("localhost:%d", s.Port),
		s.Connector(t).Address(),
	)
}
