package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/notation/fixture"
	"github.com/mark3labs/mcp-go/mcp"
	"testing"
)

func TestMarkResponseSuccessAnyOmitsRaw(t *testing.T) {
	v, e := response.SuccessAny(
		fixture.NewWrapped("a", nil, fixture.NewPrimitives("b", 0, 0, false)),
	)
	assert.FatalOnError(t, e)
	assert.String(
		t,
		"{\n\t\"Inner\": null,\n\t\"Name\": \"a\"\n}",
		v.Content[0].(mcp.TextContent).Text,
	)
}

func TestMarkResponseSuccessAnyRawKeepsRaw(t *testing.T) {
	v, e := response.SuccessAnyRaw(
		fixture.NewWrapped("a", nil, fixture.NewPrimitives("b", 0, 0, false)),
	)
	assert.FatalOnError(t, e)
	assert.String(
		t,
		"{\n\t\"Name\": \"a\",\n\t\"Inner\": null,\n\t\"Raw\": {\n\t\t\"String\": \"b\",\n\t\t\"Integer\": 0,\n\t\t\"Float\": 0,\n\t\t\"Boolean\": false\n\t}\n}",
		v.Content[0].(mcp.TextContent).Text,
	)
}
