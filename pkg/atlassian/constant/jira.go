package constant

import (
	"github.com/funtimecoding/soil/pkg/console/status/option"
	"regexp"
)

const (
	JiraDefaultProjectKeyEnvironment  = "JIRA_DEFAULT_PROJECT_KEY"
	JiraDefaultProjectNameEnvironment = "JIRA_DEFAULT_PROJECT_NAME"
	JiraDefaultIssueTypeEnvironment   = "JIRA_DEFAULT_ISSUE_TYPE"

	JiraClosedStatusEnvironment = "JIRA_CLOSED_STATUS"

	JiraTestIssueEnvironment = "JIRA_TEST_ISSUE"
	JiraTestFieldEnvironment = "JIRA_TEST_FIELD"

	JiraSearchLimit int = 100

	JiraBasicSearchLimit int = 5000

	JiraIssueCountType = "issueCount"

	JiraAllowedValuesKey = "allowedValues"

	JiraCommentPageSize = 100
	JiraCommentCap      = 20
	// API paths
	JiraBase   = "/rest/api/3"
	JiraIssue  = "/issue"
	JiraSearch = "/search/jql"

	JiraServiceDesk = "/rest/servicedeskapi"
	JiraRequest     = "/request"

	JiraDynamic = "/rest/atlassian-connect/1/app/module/dynamic"
	JiraAddon   = "/rest/atlassian-connect/1/addons"
	// Query parameter keys
	JiraMaximumResultsKey = "maxResults"
	JiraNextPageTokenKey  = "nextPageToken"
	JiraQueryKey          = "jql"
	JiraExpandKey         = "expand"
	JiraAllFields         = "*all"
	JiraChangelogExpand   = "changelog"
	JiraTimeFormat        = "2006-01-02T15:04:05.000-0700"
	// Field names
	JiraAssigneeName     = "Assignee"
	JiraAttachmentName   = "Attachment"
	JiraDescriptionName  = "Description"
	JiraIssueTypeName    = "Issue Type"
	JiraLabelsName       = "Labels"
	JiraLinkedIssuesName = "Linked Issues"
	JiraParentName       = "Parent"
	JiraProjectName      = "Project"
	JiraReporterName     = "Reporter"
	JiraSummaryName      = "Summary"
	JiraRankName         = "Rank"
	JiraFlaggedName      = "Flagged"
	JiraTeamName         = "Team"
	JiraDevelopmentName  = "Development"
	JiraParentEpic       = "parentEpic"
	// Status
	JiraInProgress = "In Progress"
	JiraClosed     = "Closed"

	JiraToDo            = "To Do"
	JiraDone            = "Done"
	ServiceDeskCanceled = "Canceled"
	ServiceDeskClosed   = "Closed"
	ServiceDeskResolved = "Resolved"
)

var (
	JiraFormat      = option.ExtendedColor.Copy()
	ServiceDeskDone = []string{
		ServiceDeskClosed,
		ServiceDeskResolved,
	}
)

const (
	JiraNoStatus      = "no status"
	JiraNoDescription = "no description"
	JiraNoLink        = "no link"
	JiraNoAge         = "no age"
	JiraNoScore       = "no score"

	JiraBlockedBy = "is blocked by"

	JiraNilValue = "nil value"

	JiraUnknownField = "unknown field"
	JiraUnknownValue = "unknown value"

	JiraDefaultPriority = "Default"
	// Issue type
	JiraBugType     = "Bug"
	JiraEpicType    = "Epic"
	JiraStoryType   = "Story"
	JiraTaskType    = "Task"
	JiraSubTaskType = "Sub-task"
)

var JiraKeyMatch = regexp.MustCompile(`[A-Z]+-\d+`)

const (
	JiraSummaryField     = "summary"
	JiraDescriptionField = "description"
)

const JiraUnknown = "unknown"
