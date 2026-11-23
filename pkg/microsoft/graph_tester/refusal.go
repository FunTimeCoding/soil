package graph_tester

import "fmt"

func Refusal(
	code string,
	directory string,
) string {
	return fmt.Sprintf(
		`{"error":"%s","error_description":"%s: The request was refused.\r\nTrace ID: 6f1c2b8e-41a0-4c55-9a3e-1b7d0c9e2f11\r\nCorrelation ID: 0d7e5a3c-88f2-4b19-a6c4-3e2f9b1d7a60\r\nTimestamp: 2026-11-23 15:42:07Z"}`,
		code,
		directory,
	)
}
