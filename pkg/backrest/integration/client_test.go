package integration

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/assert/fixture"
	"github.com/funtimecoding/soil/pkg/backrest"
	"github.com/funtimecoding/soil/pkg/backrest/face"
	"github.com/funtimecoding/soil/pkg/backrest/mock_client"
	"github.com/funtimecoding/soil/pkg/errors/command"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"testing"
)

func TestClientAndMockSatisfyFace(t *testing.T) {
	assert.Count(t, 2, []face.BackrestClient{backrest.New(), mock_client.New()})
}

func TestStanzasDecodesBackups(t *testing.T) {
	s, e := backrest.New(
		"sh",
		"-c",
		`cat "$0"`,
		fixture.Path("backrest", "backed_up.json"),
	).Stanzas(t.Context())
	assert.FatalOnError(t, e)
	assert.Count(t, 1, s)
	assert.String(t, "fixture", s[0].Name)
	assert.Integer(t, 0, s[0].Status.Code)
	assert.Count(t, 2, s[0].Backup)
	assert.String(t, "diff", s[0].Backup[1].Type)
	assert.Integer(t, int64(1794359141), s[0].Backup[1].Timestamp.Stop)
	assert.Integer(t, int64(27360975), s[0].Backup[1].Size.Database)
	assert.Integer(t, int64(3132628), s[0].Backup[1].Size.Repository.Size)
}

func TestStanzasDecodesMissingStanza(t *testing.T) {
	s, e := backrest.New(
		"sh",
		"-c",
		`cat "$0"`,
		fixture.Path("backrest", "missing_stanza.json"),
	).Stanzas(t.Context())
	assert.FatalOnError(t, e)
	assert.Integer(t, 1, s[0].Status.Code)
	assert.String(t, "missing stanza path", s[0].Status.Message)
	assert.Count(t, 0, s[0].Backup)
}

func TestStanzasUnavailableWhenCommandCannotRun(t *testing.T) {
	_, e := backrest.New("false").Stanzas(t.Context())
	assert.True(t, unavailable.Is(e))
}

func TestStanzasFailureOfPgbackrestItselfIsFault(t *testing.T) {
	_, e := backrest.New(
		"sh",
		"-c",
		`[ "$1" = version ]`,
		"sh",
	).Stanzas(t.Context())
	assert.False(t, unavailable.Is(e))
	_, okay := e.(*command.CommandError)
	assert.True(t, okay)
}

func TestReachableWhenVersionAnswers(t *testing.T) {
	assert.FatalOnError(t, backrest.New("true").Reachable(t.Context()))
}
