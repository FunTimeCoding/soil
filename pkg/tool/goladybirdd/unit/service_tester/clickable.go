package service_tester

import "github.com/funtimecoding/soil/pkg/assert"

func (o *Tester) Clickable() *Tester {
	o.t.Helper()
	o.Client.AddContext("alfa", "https://alfa.example/")
	o.Page("Probe", "Press")
	_, e := o.Service().Snapshot("")
	assert.FatalOnError(o.t, e)

	return o
}
