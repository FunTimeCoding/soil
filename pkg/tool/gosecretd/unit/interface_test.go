package unit

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gosecretd/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/gosecretd/unit/base"
	"net/http"
	"testing"
)

func TestShowOfAMissingSecretIsNotFound(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	r, e := interfaceClient(t, s).ShowSecretWithResponse(
		context.Background(),
		&client.ShowSecretParams{Path: "alfa/missing"},
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, http.StatusNotFound, r.StatusCode())
	assert.Count(t, 0, s.Reporter.Events())
}

func TestPutAndRevealOverTheInterface(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	c := interfaceClient(t, s)
	p, e := c.PutSecretWithResponse(
		context.Background(),
		&client.PutSecretParams{Path: "alfa/bravo"},
		client.PutSecretJSONRequestBody{
			Values: map[string]string{"CHARLIE": "delta-value"},
		},
	)
	assert.FatalOnError(t, e)
	assert.Strings(t, []string{"CHARLIE"}, p.JSON200.Keys)
	r, f := c.RevealSecretWithResponse(
		context.Background(),
		&client.RevealSecretParams{Path: "alfa/bravo", Key: "CHARLIE"},
	)
	assert.FatalOnError(t, f)
	assert.String(t, "delta-value", r.JSON200.Value)
}

func TestPutWithoutValuesIsRefused(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	r, e := interfaceClient(t, s).PutSecretWithResponse(
		context.Background(),
		&client.PutSecretParams{Path: "alfa/bravo"},
		client.PutSecretJSONRequestBody{Values: map[string]string{}},
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, http.StatusBadRequest, r.StatusCode())
}
