package claude_tester

import (
	"os"
	"testing"
)

func AppendFile(
	t *testing.T,
	path string,
	content string,
) {
	f, e := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if e != nil {
		t.Fatalf("open %s: %v", path, e)
	}

	if _, e = f.WriteString(content); e != nil {
		t.Fatalf("write %s: %v", path, e)
	}

	if e = f.Close(); e != nil {
		t.Fatalf("close %s: %v", path, e)
	}
}
