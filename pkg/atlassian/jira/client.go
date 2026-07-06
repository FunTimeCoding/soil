package jira

import (
	"context"
	"github.com/andygrunwald/go-jira"
	"github.com/ctreminiom/go-atlassian/v2/jira/sm"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/basic"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/field_map"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/issue/issue_enricher"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/issue/option"
)

type Client struct {
	context            context.Context
	client             *jira.Client
	locator            string
	user               string
	verbose            bool
	basic              *basic.Client
	service            *sm.Client
	fieldMap           *field_map.Map
	issueOption        *option.Issue
	enricher           *issue_enricher.Enricher
	closedStatus       []string
	defaultIssueType   string
	defaultProjectName string
	defaultProjectKey  string
}
