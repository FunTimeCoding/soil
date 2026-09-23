package scan_tester

func RootSpec() string {
	return `info:
  title: Test
paths:
  /metrics:
    get: {}
  /raw:
    get: {}
  /api/power:
    get: {}
`
}
