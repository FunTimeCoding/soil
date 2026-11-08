package face_tester

import "go/types"

func (o *Tester) Method(name string) types.Object {
	o.t.Helper()

	for i := range o.model.NumMethods() {
		if m := o.model.Method(i); m.Name() == name {
			return m
		}
	}

	o.t.Fatalf("fixture Model has no method %s", name)

	return nil
}
