package model_context

import (
	"github.com/funtimecoding/go-library/pkg/generative/model_context/parameter"
	"github.com/funtimecoding/go-library/pkg/tool/goatlassiand/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) register() {
	s.server.AddTool(
		mcp.NewTool(
			constant.JiraSearch,
			mcp.WithDescription("Search Jira issues using JQL"),
			mcp.WithString(
				parameter.Query,
				mcp.Required(),
				mcp.Description("JQL query string"),
			),
			mcp.WithNumber(
				parameter.Limit,
				mcp.Required(),
				mcp.Description("Maximum number of results (e.g. 25)"),
			),
		),
		s.searchIssues,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.JiraGetIssue,
			mcp.WithDescription("Get a Jira issue by key"),
			mcp.WithString(
				parameter.Key,
				mcp.Required(),
				mcp.Description("Issue key (e.g. PROJ-123)"),
			),
			mcp.WithBoolean(
				constant.CustomFields,
				mcp.Description(
					"Include custom field values relevant to the issue's project and type. Requires an additional API call.",
				),
			),
			mcp.WithBoolean(
				constant.Comments,
				mcp.Description(
					"Include comments on the issue.",
				),
			),
		),
		s.getIssue,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.JiraSearchProjects,
			mcp.WithDescription("Search Jira projects by key or name"),
			mcp.WithString(
				parameter.Query,
				mcp.Description("Filter by key (exact, case-insensitive) or name (case-insensitive contains). Omit to list all."),
			),
			mcp.WithNumber(
				parameter.Limit,
				mcp.Required(),
				mcp.Description("Maximum number of results (e.g. 25)"),
			),
			mcp.WithBoolean(
				constant.IncludeDescriptions,
				mcp.Description("Fetch full project details including description (slower, one API call per result). Defaults to false."),
			),
		),
		s.searchProjects,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ConfluenceSearch,
			mcp.WithDescription("Search Confluence pages using CQL or plain text"),
			mcp.WithString(
				parameter.Query,
				mcp.Required(),
				mcp.Description("CQL query string or plain text"),
			),
		),
		s.searchPages,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ConfluenceGetPage,
			mcp.WithDescription("Get a Confluence page by ID with body content as markdown"),
			mcp.WithString(
				parameter.Identifier,
				mcp.Required(),
				mcp.Description("Page ID"),
			),
			mcp.WithBoolean(
				constant.Draft,
				mcp.Description(
					"Fetch a draft page. Required for unpublished pages which return 'not found' without this flag.",
				),
			),
		),
		s.getPage,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ConfluenceGetPageDraft,
			mcp.WithDescription(
				"Get the draft layer of a Confluence page. Confluence maintains a draft layer for every page - this always returns content, even if there are no unpublished changes. Compare against the published version (confluence_get_page) to detect actual differences.",
			),
			mcp.WithString(
				parameter.Identifier,
				mcp.Required(),
				mcp.Description("Page ID"),
			),
		),
		s.getPageDraft,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ConfluenceCreatePage,
			mcp.WithDescription("Create a new Confluence page with markdown content"),
			mcp.WithString(
				constant.SpaceIdentifier,
				mcp.Required(),
				mcp.Description("Space ID"),
			),
			mcp.WithString(
				constant.ParentIdentifier,
				mcp.Required(),
				mcp.Description("Parent page ID"),
			),
			mcp.WithString(
				parameter.Title,
				mcp.Description(
					"Page title. Required for published pages, optional for drafts.",
				),
			),
			mcp.WithString(
				parameter.Body,
				mcp.Required(),
				mcp.Description("Page content in markdown"),
			),
			mcp.WithBoolean(
				constant.Draft,
				mcp.Description("Create as a draft page instead of publishing immediately."),
			),
		),
		s.createPage,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ConfluenceUpdatePage,
			mcp.WithDescription("Update a Confluence page with markdown content. Gets the current version automatically."),
			mcp.WithString(
				parameter.Identifier,
				mcp.Required(),
				mcp.Description("Page ID"),
			),
			mcp.WithString(
				parameter.Title,
				mcp.Required(),
				mcp.Description("Page title"),
			),
			mcp.WithString(
				parameter.Body,
				mcp.Required(),
				mcp.Description("Full page content in markdown"),
			),
			mcp.WithString(
				parameter.Message,
				mcp.Description("Version comment"),
			),
		),
		s.updatePage,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ConfluenceEditPage,
			mcp.WithDescription(
				"Edit a Confluence page by replacing a text fragment. Call confluence_get_page first to read the current content. The old_text must appear exactly once in the page.",
			),
			mcp.WithString(
				parameter.Identifier,
				mcp.Required(),
				mcp.Description("Page ID"),
			),
			mcp.WithString(
				constant.OldText,
				mcp.Required(),
				mcp.Description(
					"The exact text to find and replace. Must match exactly once.",
				),
			),
			mcp.WithString(
				constant.NewText,
				mcp.Required(),
				mcp.Description("The replacement text"),
			),
			mcp.WithString(
				parameter.Title,
				mcp.Description(
					"New page title. If omitted, the current title is kept.",
				),
			),
			mcp.WithString(
				parameter.Message,
				mcp.Description("Version comment"),
			),
			mcp.WithBoolean(
				constant.Draft,
				mcp.Description(
					"Edit a draft page. The page stays as a draft after editing.",
				),
			),
		),
		s.editPage,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ConfluenceListPages,
			mcp.WithDescription(
				"List pages in a Confluence space. Defaults to published pages. Set status to 'draft' to find draft pages - returns all drafts visible to the current user (the v1 API does not filter drafts by space).",
			),
			mcp.WithString(
				constant.SpaceIdentifier,
				mcp.Required(),
				mcp.Description("Space ID"),
			),
			mcp.WithString(
				constant.PageStatus,
				mcp.Description(
					"Filter by status: 'current' (default) for published pages, 'draft' for draft pages.",
				),
			),
		),
		s.listPages,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ConfluenceSetPageStatus,
			mcp.WithDescription(
				"Publish or unpublish a Confluence page. Set status to 'current' to publish a draft, or 'draft' to unpublish a published page.",
			),
			mcp.WithString(
				parameter.Identifier,
				mcp.Required(),
				mcp.Description("Page ID"),
			),
			mcp.WithString(
				constant.PageStatus,
				mcp.Required(),
				mcp.Description(
					"Target status: 'current' to publish, 'draft' to unpublish.",
				),
			),
		),
		s.setPageStatus,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ConfluenceDeletePage,
			mcp.WithDescription(
				"Delete a Confluence page. For draft pages, set draft to true - draft pages are permanently deleted, not sent to trash. Draft overlays (unpublished changes on published pages) cannot be deleted via the API.",
			),
			mcp.WithString(
				parameter.Identifier,
				mcp.Required(),
				mcp.Description("Page ID"),
			),
			mcp.WithBoolean(
				constant.Draft,
				mcp.Description(
					"Delete a draft page. Required for pages that have never been published.",
				),
			),
		),
		s.deletePage,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ConfluenceListSpaces,
			mcp.WithDescription("List all visible Confluence spaces"),
		),
		s.listSpaces,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ConfluenceGetPageChildren,
			mcp.WithDescription("List child pages of a Confluence page"),
			mcp.WithString(
				parameter.Identifier,
				mcp.Required(),
				mcp.Description("Parent page ID"),
			),
		),
		s.getPageChildren,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ConfluenceAddComment,
			mcp.WithDescription("Add a comment to a Confluence page"),
			mcp.WithString(
				parameter.Identifier,
				mcp.Required(),
				mcp.Description("Page ID"),
			),
			mcp.WithString(
				parameter.Body,
				mcp.Required(),
				mcp.Description("Comment text"),
			),
		),
		s.addPageComment,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.JiraGetCreateMeta,
			mcp.WithDescription(
				"Get required and available fields for creating a Jira issue. Returns field names, IDs, whether each field is required, its schema type, and allowed values where applicable. Call this before jira_create_issue to discover what the target project and issue type expect.",
			),
			mcp.WithString(
				constant.Project,
				mcp.Required(),
				mcp.Description("Project key (e.g. OPS)"),
			),
			mcp.WithString(
				constant.IssueType,
				mcp.Required(),
				mcp.Description(
					"Issue type name (e.g. Task, Bug, Story)",
				),
			),
			mcp.WithString(
				constant.ExpandFields,
				mcp.Description(
					"Comma-separated field names or keys to show full allowed values for. By default, fields with more than 5 allowed values are trimmed to the 5 most recent.",
				),
			),
		),
		s.getCreateMeta,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.JiraCreateIssue,
			mcp.WithDescription(
				"Create a new Jira issue. Use jira_get_create_meta first to discover required fields and allowed values for the target project and issue type.",
			),
			mcp.WithString(
				constant.Project,
				mcp.Required(),
				mcp.Description("Project key (e.g. OPS)"),
			),
			mcp.WithString(
				constant.IssueType,
				mcp.Required(),
				mcp.Description(
					"Issue type name (e.g. Task, Bug, Story)",
				),
			),
			mcp.WithString(
				constant.Summary,
				mcp.Required(),
				mcp.Description("Issue summary/title"),
			),
			mcp.WithString(
				"description",
				mcp.Description("Issue description"),
			),
			mcp.WithString(
				constant.Assignee,
				mcp.Description(
					"Assignee display name or email. Resolved to account ID automatically. Fails if no match or multiple matches.",
				),
			),
			mcp.WithString(
				constant.Labels,
				mcp.Description(
					"JSON array of label strings (e.g. [\"ZKM\",\"Infrastructure\"])",
				),
			),
			mcp.WithString(
				constant.AdditionalFields,
				mcp.Description(
					"JSON object mapping field names to values for custom or optional fields. Use jira_get_create_meta to discover available field names and allowed values.",
				),
			),
		),
		s.createIssue,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.JiraUpdateIssue,
			mcp.WithDescription(
				"Update an existing Jira issue. Only provided fields are changed. Empty strings are ignored, not applied. Returns the updated issue with a before/after diff of changed fields.",
			),
			mcp.WithString(
				parameter.Key,
				mcp.Required(),
				mcp.Description("Issue key (e.g. INF-123)"),
			),
			mcp.WithString(
				constant.Summary,
				mcp.Description("New summary/title"),
			),
			mcp.WithString(
				"description",
				mcp.Description("New description"),
			),
			mcp.WithString(
				constant.Assignee,
				mcp.Description(
					"Assignee display name or email. Resolved to account ID automatically. Pass \"none\" to unassign.",
				),
			),
			mcp.WithString(
				constant.Reporter,
				mcp.Description(
					"Reporter display name or email. Resolved to account ID automatically.",
				),
			),
			mcp.WithString(
				constant.Labels,
				mcp.Description(
					"JSON array of label strings. Replaces all existing labels.",
				),
			),
			mcp.WithString(
				constant.AdditionalFields,
				mcp.Description(
					"JSON object mapping field names to values for custom or optional fields. Use jira_get_create_meta to discover available field names and allowed values.",
				),
			),
			mcp.WithBoolean(
				constant.NoDiff,
				mcp.Description(
					"Skip the before/after diff in the response. Useful for large description changes.",
				),
			),
		),
		s.updateIssue,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.JiraGetTransitions,
			mcp.WithDescription("List available transitions for a Jira issue"),
			mcp.WithString(
				parameter.Key,
				mcp.Required(),
				mcp.Description("Issue key"),
			),
		),
		s.getTransitions,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.JiraTransitionIssue,
			mcp.WithDescription("Transition a Jira issue to a new status"),
			mcp.WithString(
				parameter.Key,
				mcp.Required(),
				mcp.Description("Issue key"),
			),
			mcp.WithString(
				constant.TransitionIdentifier,
				mcp.Required(),
				mcp.Description("Transition ID from jira_get_transitions"),
			),
		),
		s.transitionIssue,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.JiraAddComment,
			mcp.WithDescription("Add a comment to a Jira issue"),
			mcp.WithString(
				parameter.Key,
				mcp.Required(),
				mcp.Description("Issue key"),
			),
			mcp.WithString(
				parameter.Body,
				mcp.Required(),
				mcp.Description("Comment text"),
			),
		),
		s.addIssueComment,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.JiraSearchUsers,
			mcp.WithDescription(
				"Search Jira users by display name or email. Returns account IDs needed for assigning issues.",
			),
			mcp.WithString(
				parameter.Query,
				mcp.Required(),
				mcp.Description(
					"Search term (display name or email)",
				),
			),
		),
		s.searchUsers,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.JiraGetLinkTypes,
			mcp.WithDescription(
				"List available issue link types with their inward and outward labels.",
			),
		),
		s.getLinkTypes,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.JiraLinkIssues,
			mcp.WithDescription(
				"Link two Jira issues. Use jira_get_link_types to discover available link type names.",
			),
			mcp.WithString(
				parameter.Key,
				mcp.Required(),
				mcp.Description("Source issue key"),
			),
			mcp.WithString(
				constant.TargetKey,
				mcp.Required(),
				mcp.Description("Target issue key"),
			),
			mcp.WithString(
				constant.LinkType,
				mcp.Description(
					"Link type name (default: Relates). Use jira_get_link_types to see available names.",
				),
			),
		),
		s.linkIssues,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.JiraDeleteLink,
			mcp.WithDescription(
				"Delete a link between two Jira issues. Use jira_get_issue to find link IDs.",
			),
			mcp.WithString(
				parameter.Identifier,
				mcp.Required(),
				mcp.Description("Link ID"),
			),
		),
		s.deleteLink,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.JiraUpdateComment,
			mcp.WithDescription(
				"Update an existing comment on a Jira issue. Use jira_get_issue with include_comments to find comment IDs.",
			),
			mcp.WithString(
				parameter.Key,
				mcp.Required(),
				mcp.Description("Issue key"),
			),
			mcp.WithString(
				parameter.Identifier,
				mcp.Required(),
				mcp.Description("Comment ID"),
			),
			mcp.WithString(
				parameter.Body,
				mcp.Required(),
				mcp.Description("New comment body"),
			),
		),
		s.updateComment,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.JiraDeleteComment,
			mcp.WithDescription(
				"Delete a comment from a Jira issue. Use jira_get_issue with include_comments to find comment IDs.",
			),
			mcp.WithString(
				parameter.Key,
				mcp.Required(),
				mcp.Description("Issue key"),
			),
			mcp.WithString(
				parameter.Identifier,
				mcp.Required(),
				mcp.Description("Comment ID"),
			),
		),
		s.deleteComment,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.JiraGetChecklist,
			mcp.WithDescription(
				"Get the Smart Checklist items on a Jira issue. Returns indexed items with checked/unchecked status.",
			),
			mcp.WithString(
				parameter.Key,
				mcp.Required(),
				mcp.Description("Issue key"),
			),
		),
		s.getChecklist,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.JiraAddChecklistItem,
			mcp.WithDescription(
				"Add an item to a Jira issue's Smart Checklist. Added as unchecked.",
			),
			mcp.WithString(
				parameter.Key,
				mcp.Required(),
				mcp.Description("Issue key"),
			),
			mcp.WithString(
				constant.Text,
				mcp.Required(),
				mcp.Description("Checklist item text"),
			),
		),
		s.addChecklistItem,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.JiraToggleChecklistItem,
			mcp.WithDescription(
				"Toggle a Smart Checklist item between checked and unchecked.",
			),
			mcp.WithString(
				parameter.Key,
				mcp.Required(),
				mcp.Description("Issue key"),
			),
			mcp.WithNumber(
				constant.Index,
				mcp.Required(),
				mcp.Description("Item index (1-based)"),
			),
		),
		s.toggleChecklistItem,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.JiraEditChecklistItem,
			mcp.WithDescription(
				"Edit the text of a Smart Checklist item.",
			),
			mcp.WithString(
				parameter.Key,
				mcp.Required(),
				mcp.Description("Issue key"),
			),
			mcp.WithNumber(
				constant.Index,
				mcp.Required(),
				mcp.Description("Item index (1-based)"),
			),
			mcp.WithString(
				constant.Text,
				mcp.Required(),
				mcp.Description("New item text"),
			),
		),
		s.editChecklistItem,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.JiraDeleteChecklistItem,
			mcp.WithDescription(
				"Delete an item from a Jira issue's Smart Checklist.",
			),
			mcp.WithString(
				parameter.Key,
				mcp.Required(),
				mcp.Description("Issue key"),
			),
			mcp.WithNumber(
				constant.Index,
				mcp.Required(),
				mcp.Description("Item index (1-based)"),
			),
		),
		s.deleteChecklistItem,
	)
}
