package importer

import (
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store/save_option"
	"os"
	"path/filepath"
	"strings"
)

func Import(
	s *store.Store,
	directory string,
) (*Result, error) {
	entries, e := os.ReadDir(directory)

	if e != nil {
		return nil, e
	}

	existing, f := s.ListMemories("", "", false)

	if f != nil {
		return nil, f
	}

	existingHashes := map[string]bool{}

	for _, m := range existing {
		full, g := s.GetMemory(m.Identifier)

		if g != nil {
			continue
		}

		existingHashes[contentHash(full.Content)] = true
	}

	r := &Result{}

	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(
			entry.Name(),
			constant.MarkdownExtension,
		) {
			continue
		}

		if entry.Name() == "MEMORY.md" {
			continue
		}

		path := filepath.Join(directory, entry.Name())
		parsed, g := ParseFile(path)

		if g != nil {
			continue
		}

		if parsed.Content == "" {
			continue
		}

		hash := contentHash(parsed.Content)

		if existingHashes[hash] {
			r.Skipped++

			continue
		}

		o := save_option.New()
		o.Name = parsed.Name
		o.Content = parsed.Content
		o.Description = parsed.Description
		o.Type = parsed.Type
		_, e = s.CreateMemory(o)

		if e != nil {
			return nil, e
		}

		existingHashes[hash] = true
		r.Created++
	}

	return r, nil
}
