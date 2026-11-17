package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/unit/base"
	"testing"
)

func TestCreateAndGetPage(t *testing.T) {
	c := base.New(t).Tools()
	result := c.MustCallTool(
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
	getResult := c.MustCallTool(
		constant.ConfluenceGetPage,
		map[string]any{"identifier": "1"},
	)
	assert.StringContains(t, "Test Page", getResult)
}

func TestCreateDraftPage(t *testing.T) {
	c := base.New(t).Tools()
	result := c.MustCallTool(
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
	c := base.New(t).Tools()
	c.MustCallTool(
		constant.ConfluenceCreatePage,
		map[string]any{
			"space_identifier":  "1",
			"parent_identifier": "0",
			"title":             "My Draft",
			"body":              "Draft body",
			"draft":             true,
		},
	)
	result := c.MustCallTool(
		constant.ConfluenceGetPage,
		map[string]any{"identifier": "1", "draft": true},
	)
	assert.StringContains(t, "My Draft", result)
	assert.StringContains(t, "draft", result)
}

func TestPublishDraftPage(t *testing.T) {
	c := base.New(t).Tools()
	c.MustCallTool(
		constant.ConfluenceCreatePage,
		map[string]any{
			"space_identifier":  "1",
			"parent_identifier": "0",
			"title":             "Publish Me",
			"body":              "Will be published",
			"draft":             true,
		},
	)
	result := c.MustCallTool(
		constant.ConfluenceSetPageStatus,
		map[string]any{"identifier": "1", "status": "current"},
	)
	assert.StringContains(t, "current", result)
	assert.StringContains(t, "Publish Me", result)
}

func TestDraftOverlay(t *testing.T) {
	c := base.New(t).Tools()
	c.MustCallTool(
		constant.ConfluenceCreatePage,
		map[string]any{
			"space_identifier":  "1",
			"parent_identifier": "0",
			"title":             "Published Page",
			"body":              "Original content",
		},
	)
	c.MustCallTool(
		constant.ConfluenceSetPageStatus,
		map[string]any{"identifier": "1", "status": "draft"},
	)
	overlay := c.MustCallTool(
		constant.ConfluenceGetPageDraft,
		map[string]any{"identifier": "1"},
	)
	assert.StringContains(t, "draft", overlay)
	published := c.MustCallTool(
		constant.ConfluenceGetPage,
		map[string]any{"identifier": "1"},
	)
	assert.StringContains(t, "current", published)
}

func TestPublishOverlay(t *testing.T) {
	c := base.New(t).Tools()
	c.MustCallTool(
		constant.ConfluenceCreatePage,
		map[string]any{
			"space_identifier":  "1",
			"parent_identifier": "0",
			"title":             "Overlay Target",
			"body":              "Original",
		},
	)
	c.MustCallTool(
		constant.ConfluenceSetPageStatus,
		map[string]any{"identifier": "1", "status": "draft"},
	)
	result := c.MustCallTool(
		constant.ConfluenceSetPageStatus,
		map[string]any{"identifier": "1", "status": "current"},
	)
	assert.StringContains(t, "current", result)
}

func TestDeletePublishedPage(t *testing.T) {
	c := base.New(t).Tools()
	c.MustCallTool(
		constant.ConfluenceCreatePage,
		map[string]any{
			"space_identifier":  "1",
			"parent_identifier": "0",
			"title":             "Delete Me",
			"body":              "Temporary",
		},
	)
	result := c.MustCallTool(
		constant.ConfluenceDeletePage,
		map[string]any{"identifier": "1"},
	)
	assert.StringContains(t, "deleted", result)
}

func TestDeleteDraftPage(t *testing.T) {
	c := base.New(t).Tools()
	c.MustCallTool(
		constant.ConfluenceCreatePage,
		map[string]any{
			"space_identifier":  "1",
			"parent_identifier": "0",
			"title":             "Draft To Delete",
			"body":              "Will be discarded",
			"draft":             true,
		},
	)
	result := c.MustCallTool(
		constant.ConfluenceDeletePage,
		map[string]any{"identifier": "1", "draft": true},
	)
	assert.StringContains(t, "deleted", result)
}

func TestEditPage(t *testing.T) {
	c := base.New(t).Tools()
	c.MustCallTool(
		constant.ConfluenceCreatePage,
		map[string]any{
			"space_identifier":  "1",
			"parent_identifier": "0",
			"title":             "Edit Target",
			"body":              "Before edit",
		},
	)
	result := c.MustCallTool(
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
	c := base.New(t).Tools()
	c.MustCallTool(
		constant.ConfluenceCreatePage,
		map[string]any{
			"space_identifier":  "1",
			"parent_identifier": "0",
			"title":             "Edit Target",
			"body":              "Some content",
		},
	)
	result := c.MustCallToolError(
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
	c := base.New(t).Tools()
	c.MustCallTool(
		constant.ConfluenceCreatePage,
		map[string]any{
			"space_identifier":  "1",
			"parent_identifier": "0",
			"title":             "Edit Target",
			"body":              "alfa and alfa again",
		},
	)
	result := c.MustCallToolError(
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
	c := base.New(t).Tools()
	c.MustCallTool(
		constant.ConfluenceCreatePage,
		map[string]any{
			"space_identifier":  "1",
			"parent_identifier": "0",
			"title":             "Structured Page",
			"body":              "## Section One\n\n- Item A\n- Item B\n\n## Section Two\n\nParagraph here",
		},
	)
	result := c.MustCallTool(
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
	c := base.New(t).Tools()
	c.MustCallTool(
		constant.ConfluenceCreatePage,
		map[string]any{
			"space_identifier":  "1",
			"parent_identifier": "0",
			"title":             "Published",
			"body":              "Content",
		},
	)
	result := c.MustCallTool(
		constant.ConfluenceDeletePage,
		map[string]any{"identifier": "1"},
	)
	assert.StringContains(t, "deleted", result)
}

func TestDeleteDraftWithoutDraftFlagFails(t *testing.T) {
	c := base.New(t).Tools()
	c.MustCallTool(
		constant.ConfluenceCreatePage,
		map[string]any{
			"space_identifier":  "1",
			"parent_identifier": "0",
			"title":             "Draft Only",
			"body":              "Never published",
			"draft":             true,
		},
	)
	result := c.MustCallToolError(
		constant.ConfluenceDeletePage,
		map[string]any{"identifier": "1"},
	)
	assert.StringContains(t, "page not found: 1", result)
	assert.StringNotContains(t, "event_identifier", result)
}

func TestToolCount(t *testing.T) {
	c := base.New(t).Tools()
	assert.Count(t, 32, c.ListTools())
}

func TestUpdatePage(t *testing.T) {
	c := base.New(t).Tools()
	c.MustCallTool(
		constant.ConfluenceCreatePage,
		map[string]any{
			"space_identifier":  "1",
			"parent_identifier": "0",
			"title":             "Update Target",
			"body":              "Original body",
		},
	)
	result := c.MustCallTool(
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
	c := base.New(t).Tools()
	result := c.MustCallTool(
		constant.ConfluenceSearch,
		map[string]any{"query": "type=page"},
	)
	assert.StringContains(t, "[]", result)
}

func TestListSpaces(t *testing.T) {
	c := base.New(t).Tools()
	result := c.MustCallTool(constant.ConfluenceListSpaces, map[string]any{})
	assert.StringContains(t, "[]", result)
}

func TestGetPageChildren(t *testing.T) {
	c := base.New(t).Tools()
	c.MustCallTool(
		constant.ConfluenceCreatePage,
		map[string]any{
			"space_identifier":  "1",
			"parent_identifier": "0",
			"title":             "Parent",
			"body":              "Parent page",
		},
	)
	c.MustCallTool(
		constant.ConfluenceCreatePage,
		map[string]any{
			"space_identifier":  "1",
			"parent_identifier": "1",
			"title":             "Child",
			"body":              "Child page",
		},
	)
	result := c.MustCallTool(
		constant.ConfluenceGetPageChildren,
		map[string]any{"identifier": "1"},
	)
	assert.StringContains(t, "Child", result)
}

func TestAddComment(t *testing.T) {
	c := base.New(t).Tools()
	c.MustCallTool(
		constant.ConfluenceCreatePage,
		map[string]any{
			"space_identifier":  "1",
			"parent_identifier": "0",
			"title":             "Comment Target",
			"body":              "Page content",
		},
	)
	result := c.MustCallTool(
		constant.ConfluenceAddComment,
		map[string]any{"identifier": "1", "body": "A comment"},
	)
	assert.StringContains(t, "comment added", result)
}

func TestCreateDraftWithoutTitle(t *testing.T) {
	c := base.New(t).Tools()
	result := c.MustCallTool(
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
	c := base.New(t).Tools()
	c.MustCallTool(
		constant.ConfluenceCreatePage,
		map[string]any{
			"space_identifier":  "1",
			"parent_identifier": "0",
			"title":             "Old Title",
			"body":              "Content here",
		},
	)
	result := c.MustCallTool(
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
	c := base.New(t).Tools()
	c.MustCallTool(
		constant.ConfluenceCreatePage,
		map[string]any{
			"space_identifier":  "1",
			"parent_identifier": "0",
			"title":             "Message Test",
			"body":              "Before change",
		},
	)
	result := c.MustCallTool(
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
	c := base.New(t).Tools()
	c.MustCallTool(
		constant.ConfluenceCreatePage,
		map[string]any{
			"space_identifier":  "1",
			"parent_identifier": "0",
			"title":             "Draft Edit",
			"body":              "Draft before edit",
			"draft":             true,
		},
	)
	result := c.MustCallTool(
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
	c := base.New(t).Tools()
	c.MustCallTool(
		constant.ConfluenceCreatePage,
		map[string]any{
			"space_identifier":  "1",
			"parent_identifier": "0",
			"title":             "Published Page",
			"body":              "Content",
		},
	)
	c.MustCallTool(
		constant.ConfluenceCreatePage,
		map[string]any{
			"space_identifier":  "1",
			"parent_identifier": "0",
			"title":             "Draft Page",
			"body":              "Draft content",
			"draft":             true,
		},
	)
	published := c.MustCallTool(
		constant.ConfluenceListPages,
		map[string]any{"space_identifier": "1"},
	)
	assert.StringContains(t, "Published Page", published)
	drafts := c.MustCallTool(
		constant.ConfluenceListPages,
		map[string]any{"space_identifier": "1", "status": "draft"},
	)
	assert.StringContains(t, "Draft Page", drafts)
}

func TestFullDraftLifecycle(t *testing.T) {
	c := base.New(t).Tools()
	created := c.MustCallTool(
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
	read := c.MustCallTool(
		constant.ConfluenceGetPage,
		map[string]any{"identifier": "1", "draft": true},
	)
	assert.StringContains(t, "Draft content", read)
	published := c.MustCallTool(
		constant.ConfluenceSetPageStatus,
		map[string]any{"identifier": "1", "status": "current"},
	)
	assert.StringContains(t, "current", published)
	afterPublish := c.MustCallTool(
		constant.ConfluenceGetPage,
		map[string]any{"identifier": "1"},
	)
	assert.StringContains(t, "current", afterPublish)
	assert.StringContains(t, "Draft content", afterPublish)
	overlay := c.MustCallTool(
		constant.ConfluenceSetPageStatus,
		map[string]any{"identifier": "1", "status": "draft"},
	)
	assert.StringContains(t, "draft", overlay)
	stillPublished := c.MustCallTool(
		constant.ConfluenceGetPage,
		map[string]any{"identifier": "1"},
	)
	assert.StringContains(t, "current", stillPublished)
	draftOverlay := c.MustCallTool(
		constant.ConfluenceGetPageDraft,
		map[string]any{"identifier": "1"},
	)
	assert.StringContains(t, "draft", draftOverlay)
	republished := c.MustCallTool(
		constant.ConfluenceSetPageStatus,
		map[string]any{"identifier": "1", "status": "current"},
	)
	assert.StringContains(t, "current", republished)
	deleted := c.MustCallTool(
		constant.ConfluenceDeletePage,
		map[string]any{"identifier": "1"},
	)
	assert.StringContains(t, "deleted", deleted)
}
