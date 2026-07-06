//go:build local

package model_context

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/integration_test/model_context_tester"
	"testing"
)

func TestCreateAndGetPage(t *testing.T) {
	o := model_context_tester.New(t)
	defer o.Close()
	result := o.Client.MustCallTool(
		constant.ConfluenceCreatePage,
		map[string]any{
			"space_identifier":  "1",
			"parent_identifier": "0",
			"title":             "Test Page",
			"body":              "Hello world",
		},
	)
	assert.StringContains(t, "Test Page", result)
	assert.StringContains(t, "current", result)
	assert.StringContains(t, "Hello world", result)
	getResult := o.Client.MustCallTool(
		constant.ConfluenceGetPage,
		map[string]any{"identifier": "1"},
	)
	assert.StringContains(t, "Test Page", getResult)
}

func TestCreateDraftPage(t *testing.T) {
	o := model_context_tester.New(t)
	defer o.Close()
	result := o.Client.MustCallTool(
		constant.ConfluenceCreatePage,
		map[string]any{
			"space_identifier":  "1",
			"parent_identifier": "0",
			"title":             "Draft Page",
			"body":              "Draft content",
			"draft":             true,
		},
	)
	assert.StringContains(t, "Draft Page", result)
	assert.StringContains(t, "draft", result)
}

func TestGetDraftPage(t *testing.T) {
	o := model_context_tester.New(t)
	defer o.Close()
	o.Client.MustCallTool(
		constant.ConfluenceCreatePage,
		map[string]any{
			"space_identifier":  "1",
			"parent_identifier": "0",
			"title":             "My Draft",
			"body":              "Draft body",
			"draft":             true,
		},
	)
	result := o.Client.MustCallTool(
		constant.ConfluenceGetPage,
		map[string]any{"identifier": "1", "draft": true},
	)
	assert.StringContains(t, "My Draft", result)
	assert.StringContains(t, "draft", result)
}

func TestPublishDraftPage(t *testing.T) {
	o := model_context_tester.New(t)
	defer o.Close()
	o.Client.MustCallTool(
		constant.ConfluenceCreatePage,
		map[string]any{
			"space_identifier":  "1",
			"parent_identifier": "0",
			"title":             "Publish Me",
			"body":              "Will be published",
			"draft":             true,
		},
	)
	result := o.Client.MustCallTool(
		constant.ConfluenceSetPageStatus,
		map[string]any{"identifier": "1", "status": "current"},
	)
	assert.StringContains(t, "current", result)
	assert.StringContains(t, "Publish Me", result)
}

func TestDraftOverlay(t *testing.T) {
	o := model_context_tester.New(t)
	defer o.Close()
	o.Client.MustCallTool(
		constant.ConfluenceCreatePage,
		map[string]any{
			"space_identifier":  "1",
			"parent_identifier": "0",
			"title":             "Published Page",
			"body":              "Original content",
		},
	)
	o.Client.MustCallTool(
		constant.ConfluenceSetPageStatus,
		map[string]any{"identifier": "1", "status": "draft"},
	)
	overlay := o.Client.MustCallTool(
		constant.ConfluenceGetPageDraft,
		map[string]any{"identifier": "1"},
	)
	assert.StringContains(t, "draft", overlay)
	published := o.Client.MustCallTool(
		constant.ConfluenceGetPage,
		map[string]any{"identifier": "1"},
	)
	assert.StringContains(t, "current", published)
}

func TestPublishOverlay(t *testing.T) {
	o := model_context_tester.New(t)
	defer o.Close()
	o.Client.MustCallTool(
		constant.ConfluenceCreatePage,
		map[string]any{
			"space_identifier":  "1",
			"parent_identifier": "0",
			"title":             "Overlay Target",
			"body":              "Original",
		},
	)
	o.Client.MustCallTool(
		constant.ConfluenceSetPageStatus,
		map[string]any{"identifier": "1", "status": "draft"},
	)
	result := o.Client.MustCallTool(
		constant.ConfluenceSetPageStatus,
		map[string]any{"identifier": "1", "status": "current"},
	)
	assert.StringContains(t, "current", result)
}

func TestDeletePublishedPage(t *testing.T) {
	o := model_context_tester.New(t)
	defer o.Close()
	o.Client.MustCallTool(
		constant.ConfluenceCreatePage,
		map[string]any{
			"space_identifier":  "1",
			"parent_identifier": "0",
			"title":             "Delete Me",
			"body":              "Temporary",
		},
	)
	result := o.Client.MustCallTool(
		constant.ConfluenceDeletePage,
		map[string]any{"identifier": "1"},
	)
	assert.StringContains(t, "deleted", result)
}

func TestDeleteDraftPage(t *testing.T) {
	o := model_context_tester.New(t)
	defer o.Close()
	o.Client.MustCallTool(
		constant.ConfluenceCreatePage,
		map[string]any{
			"space_identifier":  "1",
			"parent_identifier": "0",
			"title":             "Draft To Delete",
			"body":              "Will be discarded",
			"draft":             true,
		},
	)
	result := o.Client.MustCallTool(
		constant.ConfluenceDeletePage,
		map[string]any{"identifier": "1", "draft": true},
	)
	assert.StringContains(t, "deleted", result)
}

func TestEditPage(t *testing.T) {
	o := model_context_tester.New(t)
	defer o.Close()
	o.Client.MustCallTool(
		constant.ConfluenceCreatePage,
		map[string]any{
			"space_identifier":  "1",
			"parent_identifier": "0",
			"title":             "Edit Target",
			"body":              "Before edit",
		},
	)
	result := o.Client.MustCallTool(
		constant.ConfluenceEditPage,
		map[string]any{
			"identifier": "1",
			"old_text":   "Before edit",
			"new_text":   "After edit",
		},
	)
	assert.StringContains(t, "After edit", result)
}

func TestEditPageNotFound(t *testing.T) {
	o := model_context_tester.New(t)
	defer o.Close()
	o.Client.MustCallTool(
		constant.ConfluenceCreatePage,
		map[string]any{
			"space_identifier":  "1",
			"parent_identifier": "0",
			"title":             "Edit Target",
			"body":              "Some content",
		},
	)
	result := o.Client.MustCallToolError(
		constant.ConfluenceEditPage,
		map[string]any{
			"identifier": "1",
			"old_text":   "nonexistent text",
			"new_text":   "replacement",
		},
	)
	assert.StringContains(t, "not found", result)
}

func TestEditPageAmbiguous(t *testing.T) {
	o := model_context_tester.New(t)
	defer o.Close()
	o.Client.MustCallTool(
		constant.ConfluenceCreatePage,
		map[string]any{
			"space_identifier":  "1",
			"parent_identifier": "0",
			"title":             "Edit Target",
			"body":              "alfa and alfa again",
		},
	)
	result := o.Client.MustCallToolError(
		constant.ConfluenceEditPage,
		map[string]any{
			"identifier": "1",
			"old_text":   "alfa",
			"new_text":   "bravo",
		},
	)
	assert.StringContains(t, "2 times", result)
}

func TestEditWithHeadingsAndLists(t *testing.T) {
	o := model_context_tester.New(t)
	defer o.Close()
	o.Client.MustCallTool(
		constant.ConfluenceCreatePage,
		map[string]any{
			"space_identifier":  "1",
			"parent_identifier": "0",
			"title":             "Structured Page",
			"body":              "## Section One\n\n- Item A\n- Item B\n\n## Section Two\n\nParagraph here",
		},
	)
	result := o.Client.MustCallTool(
		constant.ConfluenceEditPage,
		map[string]any{
			"identifier": "1",
			"old_text":   "- Item A\n- Item B",
			"new_text":   "- Item A\n- Item B\n- Item C",
		},
	)
	assert.StringContains(t, "Item C", result)
	assert.StringContains(t, "Section Two", result)
}

func TestDeletePublishedWithoutDraftFlag(t *testing.T) {
	o := model_context_tester.New(t)
	defer o.Close()
	o.Client.MustCallTool(
		constant.ConfluenceCreatePage,
		map[string]any{
			"space_identifier":  "1",
			"parent_identifier": "0",
			"title":             "Published",
			"body":              "Content",
		},
	)
	result := o.Client.MustCallTool(
		constant.ConfluenceDeletePage,
		map[string]any{"identifier": "1"},
	)
	assert.StringContains(t, "deleted", result)
}

func TestDeleteDraftWithoutDraftFlagFails(t *testing.T) {
	o := model_context_tester.New(t)
	defer o.Close()
	o.Client.MustCallTool(
		constant.ConfluenceCreatePage,
		map[string]any{
			"space_identifier":  "1",
			"parent_identifier": "0",
			"title":             "Draft Only",
			"body":              "Never published",
			"draft":             true,
		},
	)
	result := o.Client.MustCallToolError(
		constant.ConfluenceDeletePage,
		map[string]any{"identifier": "1"},
	)
	assert.StringContains(t, "not deleted", result)
}

func TestToolCount(t *testing.T) {
	o := model_context_tester.New(t)
	defer o.Close()
	assert.Count(t, 32, o.Client.ListTools())
}

func TestUpdatePage(t *testing.T) {
	o := model_context_tester.New(t)
	defer o.Close()
	o.Client.MustCallTool(
		constant.ConfluenceCreatePage,
		map[string]any{
			"space_identifier":  "1",
			"parent_identifier": "0",
			"title":             "Update Target",
			"body":              "Original body",
		},
	)
	result := o.Client.MustCallTool(
		constant.ConfluenceUpdatePage,
		map[string]any{
			"identifier": "1",
			"title":      "Updated Title",
			"body":       "New body content",
		},
	)
	assert.StringContains(t, "Updated Title", result)
	assert.StringContains(t, "New body content", result)
}

func TestSearch(t *testing.T) {
	o := model_context_tester.New(t)
	defer o.Close()
	result := o.Client.MustCallTool(
		constant.ConfluenceSearch,
		map[string]any{"query": "type=page"},
	)
	assert.StringContains(t, "null", result)
}

func TestListSpaces(t *testing.T) {
	o := model_context_tester.New(t)
	defer o.Close()
	result := o.Client.MustCallTool(
		constant.ConfluenceListSpaces,
		map[string]any{},
	)
	assert.StringContains(t, "null", result)
}

func TestGetPageChildren(t *testing.T) {
	o := model_context_tester.New(t)
	defer o.Close()
	o.Client.MustCallTool(
		constant.ConfluenceCreatePage,
		map[string]any{
			"space_identifier":  "1",
			"parent_identifier": "0",
			"title":             "Parent",
			"body":              "Parent page",
		},
	)
	o.Client.MustCallTool(
		constant.ConfluenceCreatePage,
		map[string]any{
			"space_identifier":  "1",
			"parent_identifier": "1",
			"title":             "Child",
			"body":              "Child page",
		},
	)
	result := o.Client.MustCallTool(
		constant.ConfluenceGetPageChildren,
		map[string]any{"identifier": "1"},
	)
	assert.StringContains(t, "Child", result)
}

func TestAddComment(t *testing.T) {
	o := model_context_tester.New(t)
	defer o.Close()
	o.Client.MustCallTool(
		constant.ConfluenceCreatePage,
		map[string]any{
			"space_identifier":  "1",
			"parent_identifier": "0",
			"title":             "Comment Target",
			"body":              "Page content",
		},
	)
	result := o.Client.MustCallTool(
		constant.ConfluenceAddComment,
		map[string]any{
			"identifier": "1",
			"body":       "A comment",
		},
	)
	assert.StringContains(t, "comment added", result)
}

func TestCreateDraftWithoutTitle(t *testing.T) {
	o := model_context_tester.New(t)
	defer o.Close()
	result := o.Client.MustCallTool(
		constant.ConfluenceCreatePage,
		map[string]any{
			"space_identifier":  "1",
			"parent_identifier": "0",
			"body":              "Untitled draft",
			"draft":             true,
		},
	)
	assert.StringContains(t, "draft", result)
}

func TestEditPageWithTitle(t *testing.T) {
	o := model_context_tester.New(t)
	defer o.Close()
	o.Client.MustCallTool(
		constant.ConfluenceCreatePage,
		map[string]any{
			"space_identifier":  "1",
			"parent_identifier": "0",
			"title":             "Old Title",
			"body":              "Content here",
		},
	)
	result := o.Client.MustCallTool(
		constant.ConfluenceEditPage,
		map[string]any{
			"identifier": "1",
			"old_text":   "Content here",
			"new_text":   "Updated content",
			"title":      "New Title",
		},
	)
	assert.StringContains(t, "New Title", result)
	assert.StringContains(t, "Updated content", result)
}

func TestEditPageWithMessage(t *testing.T) {
	o := model_context_tester.New(t)
	defer o.Close()
	o.Client.MustCallTool(
		constant.ConfluenceCreatePage,
		map[string]any{
			"space_identifier":  "1",
			"parent_identifier": "0",
			"title":             "Message Test",
			"body":              "Before change",
		},
	)
	result := o.Client.MustCallTool(
		constant.ConfluenceEditPage,
		map[string]any{
			"identifier": "1",
			"old_text":   "Before change",
			"new_text":   "After change",
			"message":    "Fixed typo",
		},
	)
	assert.StringContains(t, "After change", result)
	assert.StringContains(t, "Fixed typo", result)
}

func TestEditDraftPage(t *testing.T) {
	o := model_context_tester.New(t)
	defer o.Close()
	o.Client.MustCallTool(
		constant.ConfluenceCreatePage,
		map[string]any{
			"space_identifier":  "1",
			"parent_identifier": "0",
			"title":             "Draft Edit",
			"body":              "Draft before edit",
			"draft":             true,
		},
	)
	result := o.Client.MustCallTool(
		constant.ConfluenceEditPage,
		map[string]any{
			"identifier": "1",
			"old_text":   "Draft before edit",
			"new_text":   "Draft after edit",
			"draft":      true,
		},
	)
	assert.StringContains(t, "Draft after edit", result)
	assert.StringContains(t, "draft", result)
}

func TestListPages(t *testing.T) {
	o := model_context_tester.New(t)
	defer o.Close()
	o.Client.MustCallTool(
		constant.ConfluenceCreatePage,
		map[string]any{
			"space_identifier":  "1",
			"parent_identifier": "0",
			"title":             "Published Page",
			"body":              "Content",
		},
	)
	o.Client.MustCallTool(
		constant.ConfluenceCreatePage,
		map[string]any{
			"space_identifier":  "1",
			"parent_identifier": "0",
			"title":             "Draft Page",
			"body":              "Draft content",
			"draft":             true,
		},
	)
	published := o.Client.MustCallTool(
		constant.ConfluenceListPages,
		map[string]any{"space_identifier": "1"},
	)
	assert.StringContains(t, "Published Page", published)
	drafts := o.Client.MustCallTool(
		constant.ConfluenceListPages,
		map[string]any{"space_identifier": "1", "status": "draft"},
	)
	assert.StringContains(t, "Draft Page", drafts)
}

func TestFullDraftLifecycle(t *testing.T) {
	o := model_context_tester.New(t)
	defer o.Close()

	// Create draft
	created := o.Client.MustCallTool(
		constant.ConfluenceCreatePage,
		map[string]any{
			"space_identifier":  "1",
			"parent_identifier": "0",
			"title":             "Lifecycle",
			"body":              "Draft content",
			"draft":             true,
		},
	)
	assert.StringContains(t, "draft", created)

	// Read draft
	read := o.Client.MustCallTool(
		constant.ConfluenceGetPage,
		map[string]any{"identifier": "1", "draft": true},
	)
	assert.StringContains(t, "Draft content", read)

	// Publish
	published := o.Client.MustCallTool(
		constant.ConfluenceSetPageStatus,
		map[string]any{"identifier": "1", "status": "current"},
	)
	assert.StringContains(t, "current", published)

	// Read published
	afterPublish := o.Client.MustCallTool(
		constant.ConfluenceGetPage,
		map[string]any{"identifier": "1"},
	)
	assert.StringContains(t, "current", afterPublish)
	assert.StringContains(t, "Draft content", afterPublish)

	// Create overlay
	overlay := o.Client.MustCallTool(
		constant.ConfluenceSetPageStatus,
		map[string]any{"identifier": "1", "status": "draft"},
	)
	assert.StringContains(t, "draft", overlay)

	// Published version still accessible
	stillPublished := o.Client.MustCallTool(
		constant.ConfluenceGetPage,
		map[string]any{"identifier": "1"},
	)
	assert.StringContains(t, "current", stillPublished)

	// Draft accessible via get_page_draft
	draftOverlay := o.Client.MustCallTool(
		constant.ConfluenceGetPageDraft,
		map[string]any{"identifier": "1"},
	)
	assert.StringContains(t, "draft", draftOverlay)

	// Publish overlay
	republished := o.Client.MustCallTool(
		constant.ConfluenceSetPageStatus,
		map[string]any{"identifier": "1", "status": "current"},
	)
	assert.StringContains(t, "current", republished)

	// Delete
	deleted := o.Client.MustCallTool(
		constant.ConfluenceDeletePage,
		map[string]any{"identifier": "1"},
	)
	assert.StringContains(t, "deleted", deleted)
}
