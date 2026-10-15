package unit

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/mock_client"
	"net/http"
	"testing"
)

func TestVirtualMachineCustomFieldRoundTrip(t *testing.T) {
	x := context.Background()
	m := mock_client.New()
	m.SeedVirtualMachine("charlie")
	c := newClient(t, newNetboxServer(t, m, memory.New()))
	set, e := c.CreateChoiceSetWithResponse(
		x,
		client.CreateChoiceSetRequest{
			Name:    "alfa",
			Choices: []string{"echo", "foxtrot"},
		},
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, http.StatusCreated, set.StatusCode())
	field, e := c.CreateCustomFieldWithResponse(
		x,
		client.CreateCustomFieldRequest{Name: "bravo", ChoiceSet: "alfa"},
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, http.StatusCreated, field.StatusCode())
	assert.String(t, "multiselect", field.JSON201.Kind)
	updated, e := c.UpdateVirtualMachineWithResponse(
		x,
		"charlie",
		client.UpdateVirtualMachineRequest{
			CustomFields: &map[string]any{"bravo": []string{"echo", "foxtrot"}},
		},
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, http.StatusOK, updated.StatusCode())
	assert.Any(
		t,
		map[string]any{"bravo": []any{"echo", "foxtrot"}},
		*updated.JSON200.CustomFields,
	)
}

func TestAnsibleInventoryGroupsByPlaybook(t *testing.T) {
	x := context.Background()
	m := mock_client.New()
	m.SeedVirtualMachine("charlie")
	m.SeedVirtualMachine("delta")
	c := newClient(t, newNetboxServer(t, m, memory.New()))
	_, e := c.UpdateVirtualMachineWithResponse(
		x,
		"charlie",
		client.UpdateVirtualMachineRequest{
			CustomFields: &map[string]any{
				constant.AnsibleField: []string{"echo", "foxtrot"},
			},
		},
	)
	assert.FatalOnError(t, e)
	response, e := c.GetAnsibleInventoryWithResponse(
		x,
		&client.GetAnsibleInventoryParams{},
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, http.StatusOK, response.StatusCode())
	assert.String(
		t,
		`all:
    children:
        echo:
            hosts:
                charlie: {}
        foxtrot:
            hosts:
                charlie: {}
`,
		string(response.Body),
	)
}

func TestCustomFieldOnMissingChoiceSetIsNotFound(t *testing.T) {
	r := memory.New()
	c := newClient(t, newNetboxServer(t, mock_client.New(), r))
	response, e := c.CreateCustomFieldWithResponse(
		context.Background(),
		client.CreateCustomFieldRequest{Name: "bravo", ChoiceSet: "alfa"},
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, http.StatusNotFound, response.StatusCode())
	assert.Count(t, 0, r.Events())
}

func TestSetVirtualCustomFieldTool(t *testing.T) {
	m := mock_client.New()
	m.SeedVirtualMachine("charlie")
	o := newNetboxServer(t, m, memory.New()).Connect(t)
	defer o.Close()
	result := o.MustCallToolNotation(
		constant.SetVirtualCustomField,
		map[string]any{
			constant.VirtualMachine: "charlie",
			constant.Field:          "bravo",
			constant.Values:         []any{"echo"},
		},
	)
	assert.Any(
		t,
		map[string]any{"bravo": []any{"echo"}},
		result["customFields"],
	)
}
