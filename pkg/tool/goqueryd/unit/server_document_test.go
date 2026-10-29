package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"net/http"
	"testing"
)

func TestADocumentAnswersItsBody(t *testing.T) {
	answer, _ := documentAnswer(t, "notes/alfa.md")
	assert.Integer(t, http.StatusOK, answer.StatusCode())
	assert.FatalNotNil(t, answer.JSON200)
	assert.String(t, "Alfa", answer.JSON200.Title)
}

func TestAMissingDocumentAnswers404WithoutAReport(t *testing.T) {
	answer, r := documentAnswer(t, "notes/bravo.md")
	assert.Integer(t, http.StatusNotFound, answer.StatusCode())
	assert.FatalNotNil(t, answer.JSON404)
	assert.StringContains(t, "notes/bravo.md", answer.JSON404.Error)
	assert.Count(t, 0, r.Events())
}

func TestAReferenceWithoutACollectionAnswers400WithoutAReport(t *testing.T) {
	answer, r := documentAnswer(t, "alfa")
	assert.Integer(t, http.StatusBadRequest, answer.StatusCode())
	assert.FatalNotNil(t, answer.JSON400)
	assert.Count(t, 0, r.Events())
}
