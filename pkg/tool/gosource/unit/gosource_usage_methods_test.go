package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gosource"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/generated/client"
	"testing"
)

func TestUsageMethodsFilteredToAskedMethod(t *testing.T) {
	u := client.Usage{
		Type:    "Service",
		Methods: []client.UsageMethod{{Name: "Place"}, {Name: "Settle"}},
	}
	kept := gosource.UsageMethods(u, "Settle")
	assert.Count(t, 1, kept.Methods)
	assert.String(t, "Settle", kept.Methods[0].Name)
	assert.Count(t, 2, gosource.UsageMethods(u, "").Methods)
}
