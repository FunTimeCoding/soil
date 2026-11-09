package toolchain

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/generated/client"
	"testing"
)

func TestClientReadsCensus(t *testing.T) {
	c := mountedCensus(t, prepareCensus(t))
	r, e := c.GetCensus(
		t.Context(),
		&client.GetCensusParams{
			Module:  "echo",
			Package: "example/pkg/service",
			Type:    "Service",
		},
	)
	assert.FatalOnError(t, e)
	assert.Count(t, 7, r.Methods)
	assert.String(t, "Place", r.Methods[1].Name)
	assert.Strings(t, []string{"orders"}, r.Methods[1].Tables.Read)
	assert.Strings(t, []string{"sent_receipts"}, r.Methods[1].Reach.Written)
	assert.String(t, "Quiet", r.Methods[2].Name)
	assert.Count(t, 0, r.Methods[2].Uses)
}

func TestClientAnswersUnknownModuleAsNotFound(t *testing.T) {
	c := mountedCensus(t, prepareCensus(t))
	_, e := c.GetCensus(
		t.Context(),
		&client.GetCensusParams{
			Module:  "nobody",
			Package: "example/pkg/service",
			Type:    "Service",
		},
	)
	assert.True(t, not_found.Is(e))
}

func TestClientAnswersUnknownTypeAsRefused(t *testing.T) {
	c := mountedCensus(t, prepareCensus(t))
	_, e := c.GetCensus(
		t.Context(),
		&client.GetCensusParams{
			Module:  "echo",
			Package: "example/pkg/service",
			Type:    "Nothing",
		},
	)
	assert.True(t, validation.Is(e))
}
