package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/goprocess"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/client/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/generated/client"
	"testing"
)

func TestStopReachesTheProcess(t *testing.T) {
	c := mock_client.New()
	defer c.Close()
	c.ProcessList = []client.Process{{Name: "alfa", Running: true}}
	assert.FatalOnError(t, goprocess.Dispatch(c, "stop", []string{"alfa"}))
	assert.False(t, c.ProcessList[0].Running)
}

func TestAnUnknownProcessIsNotFound(t *testing.T) {
	e := goprocess.Dispatch(mock_client.New(), "restart", []string{"alfa"})
	assert.True(t, not_found.Is(e))
}

func TestAnUnknownCommandIsRefused(t *testing.T) {
	e := goprocess.Dispatch(mock_client.New(), "wipe", nil)
	assert.True(t, validation.Is(e))
}
