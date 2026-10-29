package face

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/generated/client"
)

type Client interface {
	Address() string
	Reachable(x context.Context) error
	Close()
	SearchIssues(
		x context.Context,
		query string,
		limit *int,
	) ([]*client.JiraIssue, error)
	GetIssue(
		x context.Context,
		key string,
		comments bool,
	) (*client.JiraIssue, error)
	GetTransitions(
		x context.Context,
		key string,
	) ([]*client.JiraTransition, error)
	TransitionIssue(
		x context.Context,
		key string,
		transitionIdentifier string,
	) error
	AddIssueComment(
		x context.Context,
		key string,
		body string,
	) error
	ListProjects(x context.Context) ([]*client.JiraProject, error)
	CreateIssue(
		x context.Context,
		project string,
		issueType string,
		summary string,
		description string,
		assignee string,
		labels []string,
		fields map[string]any,
	) (*client.JiraIssue, error)
	UpdateIssue(
		x context.Context,
		key string,
		summary string,
		description string,
		assignee string,
		reporter string,
		labels []string,
		fields map[string]any,
		noDiff bool,
	) (*client.IssueUpdateResult, error)
	GetCreateMeta(
		x context.Context,
		project string,
		issueType string,
		expand string,
	) ([]*client.CreateMetaField, error)
	SearchUsers(
		x context.Context,
		query string,
	) ([]*client.JiraUser, error)
	LinkIssues(
		x context.Context,
		key string,
		target string,
		linkType string,
	) error
	DeleteLink(
		x context.Context,
		identifier string,
	) error
	GetLinkTypes(x context.Context) ([]client.LinkType, error)
	UpdateComment(
		x context.Context,
		key string,
		identifier string,
		body string,
	) error
	DeleteComment(
		x context.Context,
		key string,
		identifier string,
	) error
	GetChecklist(
		x context.Context,
		key string,
	) ([]*client.ChecklistItem, error)
	AddChecklistItem(
		x context.Context,
		key string,
		text string,
	) ([]*client.ChecklistItem, error)
	EditChecklistItem(
		x context.Context,
		key string,
		index int,
		text string,
	) ([]*client.ChecklistItem, error)
	DeleteChecklistItem(
		x context.Context,
		key string,
		index int,
	) ([]*client.ChecklistItem, error)
	ToggleChecklistItem(
		x context.Context,
		key string,
		index int,
	) ([]*client.ChecklistItem, error)
	ListSpaces(x context.Context) ([]*client.ConfluenceSpace, error)
	SearchPages(
		x context.Context,
		query string,
	) ([]*client.ConfluencePage, error)
	CreatePage(
		x context.Context,
		spaceIdentifier string,
		parentIdentifier string,
		title string,
		body string,
	) (*client.ConfluencePageDetail, error)
	GetPage(
		x context.Context,
		identifier string,
	) (*client.ConfluencePageDetail, error)
	UpdatePage(
		x context.Context,
		identifier string,
		title string,
		body string,
		message *string,
	) (*client.ConfluencePageDetail, error)
	DeletePage(
		x context.Context,
		identifier string,
		draft bool,
	) error
	GetPageChildren(
		x context.Context,
		identifier string,
	) ([]*client.ConfluencePage, error)
	AddPageComment(
		x context.Context,
		identifier string,
		body string,
	) error
	GetPageDraft(
		x context.Context,
		identifier string,
	) (*client.ConfluencePageDetail, error)
	SetPageStatus(
		x context.Context,
		identifier string,
		status string,
	) (*client.ConfluencePageDetail, error)
	EditPage(
		x context.Context,
		identifier string,
		oldText string,
		newText string,
		title string,
		message string,
		draft bool,
	) (*client.ConfluencePageDetail, error)
	ListPages(
		x context.Context,
		space string,
		status string,
	) ([]*client.ConfluencePage, error)
}
