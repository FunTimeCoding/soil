package security_tester

import "github.com/funtimecoding/soil/pkg/tool/gosecurityd/service/remediation"

func (o *Tester) Remediation() *remediation.Service {
	o.t.Helper()

	return remediation.New(o.System(), o.Flag())
}
