package unit

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/client/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/generated/client"
	"testing"
)

func TestAWriteWithoutAResultPrintsNothingOnSuccess(t *testing.T) {
	for _, arguments := range writes() {
		c := mock_client.New()
		var output, failure string
		var code int
		stray := assert.Capture(
			t,
			func() { output, failure, code = run(t, c, arguments...) },
		)
		assert.Integer(t, 0, code)
		assert.String(t, "", output)
		assert.String(t, "", failure)
		assert.String(t, "", stray)
		c.Close()
	}
}

func TestAFailedWriteExitsWithItsMessage(t *testing.T) {
	for _, arguments := range writes() {
		c := mock_client.New()
		c.Failure = errors.New("foxtrot")
		_, failure, code := run(t, c, arguments...)
		assert.Integer(t, 1, code)
		assert.StringContains(t, "foxtrot", failure)
		c.Close()
	}
}

func TestSearchIssuesEmitsTheIssues(t *testing.T) {
	c := mock_client.New()
	defer c.Close()
	c.Found = []*client.JiraIssue{{Key: "ALFA-1"}}
	output, _, code := run(t, c, "search-issues", "bravo")
	assert.Integer(t, 0, code)
	var found []client.JiraIssue
	notation.MustDecode(output, &found, false)
	assert.Count(t, 1, found)
	assert.String(t, "ALFA-1", found[0].Key)
}
