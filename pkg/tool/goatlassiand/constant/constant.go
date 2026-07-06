package constant

import "github.com/funtimecoding/soil/pkg/identity"

var Identity = identity.New(
	"goatlassiand",
	"Jira and Confluence bridge",
	"goatlassiand",
).WithInstructions(
	"Jira and Confluence - search and manage issues, transitions, checklists, comments, issue links. Confluence: search, read, create, update, edit, and delete pages. Supports draft pages - create as draft, read drafts and unpublished changes, publish and unpublish. Use confluence_edit_page for partial edits (old_text/new_text replacement) instead of rewriting the full page body.",
)

const (
	JiraSearch              = "jira_search"
	JiraGetIssue            = "jira_get_issue"
	JiraSearchProjects      = "jira_search_projects"
	JiraGetTransitions      = "jira_get_transitions"
	JiraTransitionIssue     = "jira_transition_issue"
	JiraAddComment          = "jira_add_comment"
	JiraGetCreateMeta       = "jira_get_create_meta"
	JiraCreateIssue         = "jira_create_issue"
	JiraUpdateIssue         = "jira_update_issue"
	JiraSearchUsers         = "jira_search_users"
	JiraGetChecklist        = "jira_get_checklist"
	JiraAddChecklistItem    = "jira_add_checklist_item"
	JiraToggleChecklistItem = "jira_toggle_checklist_item"
	JiraEditChecklistItem   = "jira_edit_checklist_item"
	JiraDeleteChecklistItem = "jira_delete_checklist_item"
	JiraGetLinkTypes        = "jira_get_link_types"
	JiraLinkIssues          = "jira_link_issues"
	JiraDeleteLink          = "jira_delete_link"
	JiraUpdateComment       = "jira_update_comment"
	JiraDeleteComment       = "jira_delete_comment"

	ConfluenceSearch          = "confluence_search"
	ConfluenceGetPage         = "confluence_get_page"
	ConfluenceCreatePage      = "confluence_create_page"
	ConfluenceUpdatePage      = "confluence_update_page"
	ConfluenceListSpaces      = "confluence_list_spaces"
	ConfluenceGetPageChildren = "confluence_get_page_children"
	ConfluenceDeletePage      = "confluence_delete_page"
	ConfluenceListPages       = "confluence_list_pages"
	ConfluenceGetPageDraft    = "confluence_get_page_draft"
	ConfluenceEditPage        = "confluence_edit_page"
	ConfluenceSetPageStatus   = "confluence_set_page_status"
	ConfluenceAddComment      = "confluence_add_comment"
	AllowedValueLimit         = 5

	ChecklistField       = "Checklists"
	Unassign             = "none"
	AdditionalFields     = "additional_fields"
	Assignee             = "assignee"
	Comments             = "include_comments"
	CustomFields         = "include_custom_fields"
	Draft                = "draft"
	ExpandFields         = "expand_fields"
	IncludeDescriptions  = "include_descriptions"
	IssueType            = "issue_type"
	Labels               = "labels"
	NoDiff               = "no_diff"
	NewText              = "new_text"
	OldText              = "old_text"
	PageStatus           = "status"
	Index                = "index"
	LinkType             = "link_type"
	Text                 = "text"
	ParentIdentifier     = "parent_identifier"
	Project              = "project"
	Reporter             = "reporter"
	TargetKey            = "target_key"
	SpaceIdentifier      = "space_identifier"
	Summary              = "summary"
	TransitionIdentifier = "transition_identifier"
)
