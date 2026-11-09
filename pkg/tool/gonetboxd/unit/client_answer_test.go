package unit

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/reported"
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/unit/client_tester"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/unit/refusing_source"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/unit/service_tester"
	"testing"
)

func TestClientReadsDaemon(t *testing.T) {
	m := service_tester.NewSource(t)
	m.SeedDevice("alfa")
	c := client_tester.New(t, m)
	v, e := c.GetDevice(context.Background(), "alfa")
	assert.FatalOnError(t, e)
	assert.String(t, "alfa", v.Name)
	assert.FatalOnError(t, c.Reachable(context.Background()))
}

func TestClientAnswersUnknownDeviceAsNotFound(t *testing.T) {
	_, e := client_tester.New(t, service_tester.NewSource(t)).GetDevice(
		context.Background(),
		"alfa",
	)
	assert.True(t, not_found.Is(e))
}

func TestClientAnswersRejectedCreateAsValidation(t *testing.T) {
	_, e := client_tester.New(
		t,
		refusing_source.New(validation.New("slug: exists")),
	).CreateSite(context.Background(), "alfa")
	assert.True(t, validation.Is(e))
	assert.StringContains(t, "slug: exists", e.Error())
}

func TestClientAnswersRefusedTokenAsReported(t *testing.T) {
	_, e := client_tester.New(
		t,
		refusing_source.New(
			unexpected.Format("netbox status 403: no permission"),
		),
	).CreateSite(context.Background(), "alfa")
	assert.True(t, reported.Is(e))
	assert.StringContains(t, "netbox status 403: no permission", e.Error())
}
