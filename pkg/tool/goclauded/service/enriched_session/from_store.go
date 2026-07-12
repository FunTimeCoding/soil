package enriched_session

import "github.com/funtimecoding/soil/pkg/tool/goclauded/store/session"

func FromStore(i session.Session, active bool) *Session {
	result := New()
	result.Identifier = i.Identifier
	result.Slug = i.Slug
	result.Timestamp = i.SessionTimestamp
	result.WorkDirectory = i.WorkDirectory
	result.Branch = i.Branch
	result.Lines = i.Lines
	result.Name = i.Name
	result.Alias = i.AliasValue()
	result.Description = i.Description
	result.TurnCount = i.TurnCount
	result.Active = active

	return result
}
