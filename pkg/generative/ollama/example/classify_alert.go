package example

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/console"
	generative "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/generative/ollama"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/system/constant"
	"github.com/funtimecoding/soil/pkg/system/join"
)

func ClassifyAlert() {
	o := ollama.NewEnvironment()

	if false {
		alert := "ECCMemoryError"
		// DiskNearFull answer is consistently "not broken"
		// ECCMemoryError answer is inconsistent
		r := o.GenerateNotation(
			fmt.Sprintf(
				"Answer a JSON object with 2 strings: Reason and Answer. Assess in one short sentence first, then answer already-broken or not-yet-broken. Does this Prometheus alert indicate something is already-broken or not-yet-broken: %s\nAlready-broken examples: DiskFull, Timeout\nNot-yet-broken examples: DiskNearFull, HighLatency",
				alert,
			),
		)
		console.Line(r.Text)
		r.Print()
	}

	if true {
		p := ollama.ClassifyAlert()
		r := o.GenerateNotation(p.Render())
		console.Format("Response: %+v\n", r)
		response := p.ParseResponse(r.Text)
		console.Format("To save: %s", response)
		base := join.Absolute(constant.Temporary, generative.OllamaScheme)
		system.EnsurePathExists(base)
		system.SaveFile(
			join.Absolute(
				base,
				fmt.Sprintf("response-%d.json", len(system.Files(base))+1),
			),
			response,
		)
	}
}
