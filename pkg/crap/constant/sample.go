package constant

const FixtureGremlins = `{
  "go_module": "example.test/m",
  "files": [
    {
      "file_name": "pkg/a/half.go",
      "mutations": [
        {"type": "CONDITIONALS_BOUNDARY", "status": "LIVED", "line": 4},
        {"type": "ARITHMETIC", "status": "KILLED", "line": 5}
      ]
    },
    {
      "file_name": "pkg/a/covered.go",
      "mutations": [
        {"type": "CONDITIONALS_NEGATION", "status": "KILLED", "line": 4},
        {"type": "CONTROL_FLOW", "status": "NOT COVERED", "line": 99}
      ]
    }
  ]
}`
