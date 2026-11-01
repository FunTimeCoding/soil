package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/embed_counter"
	"path/filepath"
	"testing"
)

func TestAnUnsetTokenizerNamesTheFetch(t *testing.T) {
	t.Setenv(constant.EmbedTokenizerEnvironment, "")
	_, e := embed_counter.NewEnvironment()
	assert.StringContains(t, "EMBED_TOKENIZER not set", e.Error())
	assert.StringContains(
		t,
		"curl -L --create-dirs -o ~/.local/share/soil/nomic-embed-text/tokenizer.json https://huggingface.co/nomic-ai/nomic-embed-text-v1.5/resolve/main/tokenizer.json",
		e.Error(),
	)
}

func TestAMissingTokenizerNamesTheFetchToItsPath(t *testing.T) {
	path := filepath.Join(t.TempDir(), constant.RerankTokenizerFile)
	t.Setenv(constant.EmbedTokenizerEnvironment, path)
	_, e := embed_counter.NewEnvironment()
	assert.StringContains(t, "load embed tokenizer", e.Error())
	assert.StringContains(
		t,
		join.Empty("curl -L --create-dirs -o ", path, " "),
		e.Error(),
	)
}
