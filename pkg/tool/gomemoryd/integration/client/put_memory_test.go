package client

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/integration/base"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store/save_option"
	"github.com/funtimecoding/soil/pkg/web"
	"net/http"
	"testing"
)

func TestPutMemoryWithDescriptionOnly(t *testing.T) {
	s := base.New(t)
	o := save_option.New()
	o.Name = "pace"
	o.Content = "Original body."
	o.Description = "a test"
	identifier, e := s.Store().CreateMemory(o)
	assert.FatalOnError(t, e)
	c, f := client.NewClientWithResponses(
		fmt.Sprintf("http://localhost:%d", s.Port),
		client.WithRequestEditorFn(
			web.BearerEditor(constant.ModelContextTestToken),
		),
	)
	errors.PanicOnError(f)
	r, g := c.PutMemoryWithResponse(
		context.Background(),
		identifier,
		client.PutMemoryJSONRequestBody{Description: new("a better test")},
	)
	assert.FatalOnError(t, g)
	assert.Integer(t, http.StatusOK, r.StatusCode())
	m, h := s.Store().GetMemory(identifier)
	assert.FatalOnError(t, h)
	assert.String(t, "pace", m.Name)
	assert.String(t, "Original body.", m.Content)
	assert.String(t, "a better test", m.Description)
}

func TestPutMemoryWithEmptyContentSetsIt(t *testing.T) {
	s := base.New(t)
	o := save_option.New()
	o.Name = "pace"
	o.Content = "Original body."
	o.Description = "a test"
	identifier, e := s.Store().CreateMemory(o)
	assert.FatalOnError(t, e)
	c, f := client.NewClientWithResponses(
		fmt.Sprintf("http://localhost:%d", s.Port),
		client.WithRequestEditorFn(
			web.BearerEditor(constant.ModelContextTestToken),
		),
	)
	errors.PanicOnError(f)
	r, g := c.PutMemoryWithResponse(
		context.Background(),
		identifier,
		client.PutMemoryJSONRequestBody{Content: new("")},
	)
	assert.FatalOnError(t, g)
	assert.Integer(t, http.StatusOK, r.StatusCode())
	m, h := s.Store().GetMemory(identifier)
	assert.FatalOnError(t, h)
	assert.String(t, "pace", m.Name)
	assert.String(t, "", m.Content)
	assert.String(t, "a test", m.Description)
}
