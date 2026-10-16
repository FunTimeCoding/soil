package constant

const FixtureInstance = `
name: operated
alertmanager:
  host: alertmanager-operated.observability
  port: 9093
  link:
    host: a.example
    secure: true
prometheus:
  host: prometheus-operated.observability
  port: 9090
`
