package unit

import "github.com/funtimecoding/soil/pkg/strings/join"

func searchTranscript() string {
	return join.Empty(
		searchLine(
			"u1",
			"user",
			"2026-10-07T20:00:00Z",
			false,
			`"how does the legacy protocol pin work"`,
		),
		searchLine(
			"a1",
			"assistant",
			"2026-10-07T20:01:00Z",
			false,
			`[{"type":"text","text":"advertise legacy versions through the discover RPC"}]`,
		),
		searchLine(
			"a2",
			"assistant",
			"2026-10-07T20:02:00Z",
			false,
			`[{"type":"tool_use","id":"t1","name":"Edit","input":{"file_path":"/tmp/setup.go","old_string":"heartbeat","new_string":"WithStreamableHTTPProtocolVersions"}}]`,
		),
		searchLine(
			"a3",
			"assistant",
			"2026-10-07T20:03:00Z",
			false,
			`[{"type":"tool_use","id":"t2","name":"Bash","input":{"command":"go test ./...","description":"run the tests"}}]`,
		),
		searchLine(
			"r1",
			"user",
			"2026-10-07T20:04:00Z",
			false,
			`[{"type":"tool_result","tool_use_id":"t2","content":"PASS legacy suite"}]`,
		),
		searchLine(
			"m1",
			"user",
			"2026-10-07T20:05:00Z",
			false,
			`"<command-name>/clear</command-name>"`,
		),
		searchLine(
			"m2",
			"user",
			"2026-10-07T20:06:00Z",
			true,
			`"Caveat: hidden meta line"`,
		),
	)
}
