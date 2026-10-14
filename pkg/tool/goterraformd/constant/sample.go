package constant

const (
	FixtureLockDetail = `{"ID":"11111111-2222-4333-8444-555555555555","Operation":"OperationTypePlan","Info":"","Who":"user@host.example","Version":"1.0.0","Created":"2020-01-02T03:04:05.123456789Z","Path":""}`

	FixtureLockTemplate = `{"ID":%q,"Operation":"OperationTypeApply","Info":"","Who":%q,"Version":"1.16.5","Created":%q,"Path":""}`

	FixtureRecorder = `#!/bin/sh
echo "$*" >> %q
[ "$1" = force-unlock ] && exit %d
exit 0
`
)
