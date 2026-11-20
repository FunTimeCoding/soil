package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/unit/document_tester"
	"net/http"
	"testing"
)

func TestDocumentAnswersItsBody(t *testing.T) {
	answer, _ := document_tester.New(t).Answer("notes/alfa.md")
	assert.Integer(t, http.StatusOK, answer.StatusCode())
	assert.FatalNotNil(t, answer.JSON200)
	assert.String(t, "Alfa", answer.JSON200.Title)
}

func TestMissingDocumentAnswers404WithoutReport(t *testing.T) {
	answer, r := document_tester.New(t).Answer("notes/bravo.md")
	assert.Integer(t, http.StatusNotFound, answer.StatusCode())
	assert.FatalNotNil(t, answer.JSON404)
	assert.StringContains(t, "notes/bravo.md", answer.JSON404.Error)
	assert.Count(t, 0, r.Events())
}

func TestReferenceWithoutCollectionAnswers400WithoutReport(t *testing.T) {
	answer, r := document_tester.New(t).Answer("alfa")
	assert.Integer(t, http.StatusBadRequest, answer.StatusCode())
	assert.FatalNotNil(t, answer.JSON400)
	assert.Count(t, 0, r.Events())
}
