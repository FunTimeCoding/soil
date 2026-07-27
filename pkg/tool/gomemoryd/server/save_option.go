package server

import (
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store/save_option"
)

func saveOption(r server.MemoryRequest) *save_option.Option {
	o := save_option.New()
	o.Name = r.Name
	o.Content = r.Content
	o.Description = r.Description
	o.ParentIdentifier = r.ParentIdentifier

	if r.Type != nil {
		o.Type = *r.Type
	}

	if r.Scope != nil {
		o.Scope = *r.Scope
	}

	if r.Tags != nil {
		o.Tags = *r.Tags
	}

	if r.Metadata != nil {
		o.Metadata = *r.Metadata
	}

	if r.Source != nil {
		o.Source = *r.Source
	}

	if r.ProvenanceFile != nil {
		o.ProvenanceFile = *r.ProvenanceFile
	}

	if r.ProvenanceAnchor != nil {
		o.ProvenanceAnchor = *r.ProvenanceAnchor
	}

	if r.ProvenanceHash != nil {
		o.ProvenanceHash = *r.ProvenanceHash
	}

	if r.Ordinal != nil {
		o.Ordinal = *r.Ordinal
	}

	return o
}
