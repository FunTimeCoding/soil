package unit

import "fmt"

func transcriptLine(
	name string,
	timestamp string,
) string {
	return fmt.Sprintf(
		`{"type":"assistant","timestamp":"%s","message":{"content":[{"type":"tool_use","id":"t1","name":"%s","input":{}}]}}%s`,
		timestamp,
		name,
		"\n",
	)
}
