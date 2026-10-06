package chunk

import "github.com/funtimecoding/soil/pkg/face"

func chunkText(
	content string,
	c face.TokenCounter,
) []Chunk {
	tables := findTables(content)
	pulled := map[int]bool{}

	for i, t := range tables {
		if len(t.rows) > 0 && c.Count(content[t.start:t.end]) > c.Allowance() {
			pulled[i] = true
		}
	}

	for {
		result := assemble(content, tables, pulled, c)
		changed := false

		for _, k := range result {
			if k.Length != len(k.Text) || c.Count(k.Text) <= c.Allowance() {
				continue
			}

			for i, t := range tables {
				if pulled[i] || len(t.rows) == 0 {
					continue
				}

				if t.start < k.Position+k.Length && t.end > k.Position {
					pulled[i] = true
					changed = true
				}
			}
		}

		if !changed {
			return result
		}
	}
}
