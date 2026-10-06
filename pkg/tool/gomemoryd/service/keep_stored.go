package service

import (
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store/record"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store/save_option"
)

func keepStored(
	o *save_option.Option,
	existing *record.Memory,
) {
	if o.Tags == nil {
		o.Tags = existing.Tags
	}

	if o.Metadata == nil {
		o.Metadata = existing.Metadata
	}

	o.Metadata = applyBase(o.Metadata, o.Base)

	if o.ProvenanceHash == "" {
		o.ProvenanceHash = existing.ProvenanceHash
	}

	if o.Ordinal == 0 {
		o.Ordinal = existing.Ordinal
	}
}
