package flow

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gogated/integration/tester"
	"testing"
)

func TestRegistrationReturnsClientCredentials(t *testing.T) {
	o := tester.New(t)
	result := o.Register(t, []string{"http://localhost/callback"})
	assert.StringContains(t, "-", result.ClientIdentifier)
	assert.StringContains(t, "-", result.ClientSecret)
	assert.Integer(t, 1, len(result.RedirectLocators))
	assert.String(t, "http://localhost/callback", result.RedirectLocators[0])
}

func TestRegistrationCreatesUniqueClients(t *testing.T) {
	o := tester.New(t)
	first := o.Register(t, []string{"http://localhost/a"})
	second := o.Register(t, []string{"http://localhost/b"})

	if first.ClientIdentifier == second.ClientIdentifier {
		t.Fatal("expected different client identifiers")
	}
}
