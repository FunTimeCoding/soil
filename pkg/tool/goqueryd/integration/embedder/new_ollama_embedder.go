package embedder

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"strings"
	"testing"
)

func newOllamaEmbedder(t *testing.T) face.ReachableEmbedder {
	t.Helper()
	target := environment.Optional(constant.EmbedTargetEnvironment)

	if !strings.HasPrefix(target, join.Empty(constant.OllamaScheme, ":")) {
		t.Skipf(
			"the claim is the ollama client's truncate flag; target is %q",
			target,
		)
	}

	return newEmbedder(t)
}
