package markup_tester

import "github.com/andybalholm/cascadia"

func (o *Tester) matcher(selector string) cascadia.Selector {
	o.t.Helper()
	m, e := cascadia.Compile(selector)

	if e != nil {
		o.t.Fatalf("selector %q does not compile: %v", selector, e)
	}

	return m
}
