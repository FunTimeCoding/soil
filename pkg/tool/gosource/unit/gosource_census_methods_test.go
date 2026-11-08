package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gosource"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/generated/client"
	"testing"
)

func TestTheCensusKeepsOnlyTheAskedMethod(t *testing.T) {
	c := client.Census{
		Type:    "Service",
		Methods: []client.CensusMethod{{Name: "Place"}, {Name: "Settle"}},
	}
	kept := gosource.CensusMethods(c, "Settle")
	assert.Count(t, 1, kept.Methods)
	assert.String(t, "Settle", kept.Methods[0].Name)
	assert.Count(t, 2, gosource.CensusMethods(c, "").Methods)
}
