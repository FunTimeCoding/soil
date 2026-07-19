package unit_test

import (
	"github.com/andygrunwald/go-jira"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/field_map"
	"testing"
)

func TestJiraFieldMap(t *testing.T) {
	assert.NotNil(t, field_map.New([]jira.Field{}))
}
