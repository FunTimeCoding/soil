package unit_test

import (
	"github.com/funtimecoding/soil/pkg/text/dictionary"
	"os"
	"path/filepath"
	"testing"
)

func TestMerge(t *testing.T) {
	d := t.TempDir()
	target := filepath.Join(d, "target.dic")
	source := filepath.Join(d, "source.dic")

	if e := os.WriteFile(
		target,
		[]byte("# English\nalfa\ncharlie\n\n# Go\ngoroutine\n"),
		0o644,
	); e != nil {
		t.Fatal(e)
	}

	if e := os.WriteFile(
		source,
		[]byte("# English\nbravo\ncharlie\n\n# Kubernetes\nkubectl\n"),
		0o644,
	); e != nil {
		t.Fatal(e)
	}

	added := dictionary.Merge(target, source)

	if added != 2 {
		t.Fatalf("expected 2 added, got %d", added)
	}

	categories := dictionary.Read(target)

	if len(categories) != 3 {
		t.Fatalf("expected 3 categories, got %d", len(categories))
	}

	english := categories[0]

	if english.Name != "English" {
		t.Fatalf("expected English, got %s", english.Name)
	}

	if len(english.Words) != 3 {
		t.Fatalf(
			"expected 3 words in English, got %d",
			len(english.Words),
		)
	}

	if english.Words[0] != "alfa" ||
		english.Words[1] != "bravo" ||
		english.Words[2] != "charlie" {
		t.Fatalf("unexpected words: %v", english.Words)
	}

	if categories[2].Name != "Kubernetes" {
		t.Fatalf(
			"expected Kubernetes category, got %s",
			categories[2].Name,
		)
	}
}
