package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecretd/client"
	"github.com/funtimecoding/soil/pkg/tool/gosecretd/client/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/gosecretd/face"
	generated "github.com/funtimecoding/soil/pkg/tool/gosecretd/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/gosecretd/unit/base"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"testing"
)

func TestMockAnswersUnknownSecretAsDaemonDoes(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	c := client.New(
		locator.New("localhost").Port(s.ContextServer.Port).Insecure(),
		generativeConstant.ModelContextTestToken,
	)
	defer c.Close()
	m := mock_client.New()
	defer m.Close()
	x := t.Context()

	for _, f := range []face.Client{c, m} {
		_, e := f.PutSecret(
			x,
			&generated.PutSecretParams{Path: "alfa/bravo"},
			generated.PutSecretJSONRequestBody{
				Values: map[string]string{"CHARLIE": "delta-value"},
			},
		)
		assert.FatalOnError(t, e)
		_, g := f.ShowSecret(
			x,
			&generated.ShowSecretParams{Path: "alfa/missing"},
		)
		assert.True(t, not_found.Is(g))
		_, h := f.RevealSecret(
			x,
			&generated.RevealSecretParams{Path: "alfa/missing", Key: "CHARLIE"},
		)
		assert.True(t, not_found.Is(h))
		_, i := f.RevealSecret(
			x,
			&generated.RevealSecretParams{Path: "alfa/bravo", Key: "ECHO"},
		)
		assert.True(t, not_found.Is(i))
	}
}
