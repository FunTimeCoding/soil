package project_match

import "github.com/andygrunwald/go-jira"

type Match struct {
	Key    string
	Name   string
	Type   string
	Detail *jira.Project
}
