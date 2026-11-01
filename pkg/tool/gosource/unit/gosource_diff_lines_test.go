package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gosource"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/generated/client"
	"testing"
)

func TestDiffLinesPrintOneLinePerEntry(t *testing.T) {
	assert.Strings(
		t,
		[]string{
			"counts unchanged 40 moved 1 removed 1",
			"module lima v0.1.0 v0.2.0",
			"module mike v0.3.0 staged",
			"unparsed to broken.go",
			"moved bravo bravo New (bravo/new.go:5) <- alfa alfa New (alfa/new.go:5)",
			"removed . alfa Echo (alfa.go:3)",
		},
		gosource.DiffLines(
			client.DiffResponse{
				Counts: []client.ClassCount{
					{Class: "unchanged", Count: 40},
					{Class: "moved", Count: 1},
					{Class: "removed", Count: 1},
				},
				Modules: []client.ModuleRange{
					{Name: "lima", From: "v0.1.0", To: "v0.2.0"},
					{Name: "mike", From: "v0.3.0", To: ""},
				},
				Entries: []client.Entry{
					{
						Class: "moved",
						From: &client.Location{
							Key:  "alfa alfa New",
							File: "alfa/new.go",
							Line: 5,
						},
						To: &client.Location{
							Key:  "bravo bravo New",
							File: "bravo/new.go",
							Line: 5,
						},
					},
					{
						Class: "removed",
						From: &client.Location{
							Key:  ". alfa Echo",
							File: "alfa.go",
							Line: 3,
						},
					},
				},
				ToUnparsed: []string{"broken.go"},
			},
		),
	)
}
