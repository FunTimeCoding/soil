package constant

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/console/status/option"
)

const (
	DefaultProjectKeyEnvironment  = "JIRA_DEFAULT_PROJECT_KEY"
	DefaultProjectNameEnvironment = "JIRA_DEFAULT_PROJECT_NAME"
	DefaultIssueTypeEnvironment   = "JIRA_DEFAULT_ISSUE_TYPE"

	ClosedStatusEnvironment = "JIRA_CLOSED_STATUS"

	TestIssueEnvironment = "JIRA_TEST_ISSUE"
	TestFieldEnvironment = "JIRA_TEST_FIELD"

	SearchLimit int = 100

	BasicSearchLimit int = 5000

	IssueCountType = "issueCount"

	AllowedValuesKey = "allowedValues"

	CommentPageSize = 100
	CommentCap      = 20
)

// API paths
const (
	Base   = "/rest/api/3"
	Issue  = "/issue"
	Search = "/search/jql"

	ServiceDesk = "/rest/servicedeskapi"
	Request     = "/request"

	Dynamic = "/rest/atlassian-connect/1/app/module/dynamic"
	Addon   = "/rest/atlassian-connect/1/addons"
)

// Query parameter keys
const (
	MaximumResultsKey = "maxResults"
	NextPageTokenKey  = "nextPageToken"
	QueryKey          = "jql"
	ExpandKey         = "expand"
)

const AllFields = "*all" // Query parameter value
const ChangelogExpand = "changelog"
const TimeFormat = "2006-01-02T15:04:05.000-0700"

// Field names
const (
	AssigneeName     = "Assignee"
	AttachmentName   = "Attachment"
	DescriptionName  = "Description"
	IssueTypeName    = "Issue Type"
	LabelsName       = "Labels"
	LinkedIssuesName = "Linked Issues"
	ParentName       = "Parent"
	ProjectName      = "Project"
	ReporterName     = "Reporter"
	SummaryName      = "Summary"
	RankName         = "Rank"
	FlaggedName      = "Flagged"
	TeamName         = "Team"
	DevelopmentName  = "Development"
)

const ParentEpic = "parentEpic" // Field

// Status
const (
	InProgress = "In Progress"
	Closed     = "Closed"

	ToDo = "To Do"
	Done = "Done"
)

var (
	ErrorNotFound = errors.New("not found")
	Format        = option.ExtendedColor.Copy()
)

const (
	ServiceDeskCanceled = "Canceled"
	ServiceDeskClosed   = "Closed"
	ServiceDeskResolved = "Resolved"
)

var ServiceDeskDone = []string{
	ServiceDeskClosed,
	ServiceDeskResolved,
}
