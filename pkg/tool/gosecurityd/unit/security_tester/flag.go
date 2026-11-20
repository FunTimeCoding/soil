package security_tester

import "github.com/funtimecoding/soil/pkg/tool/gosecurityd/service/flag"

func (o *Tester) Flag() *flag.Service {
	o.t.Helper()

	if o.flag != nil {
		return o.flag
	}

	o.flag = flag.New(o.Store(), o.Vulnerability())

	return o.flag
}
