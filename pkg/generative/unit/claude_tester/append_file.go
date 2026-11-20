package claude_tester

import "os"

func (o *Tester) AppendFile(
	path string,
	content string,
) {
	o.t.Helper()
	f, e := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if e != nil {
		o.t.Fatalf("open %s: %v", path, e)
	}

	if _, e = f.WriteString(content); e != nil {
		o.t.Fatalf("write %s: %v", path, e)
	}

	if e = f.Close(); e != nil {
		o.t.Fatalf("close %s: %v", path, e)
	}
}
