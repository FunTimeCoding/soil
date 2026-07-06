package constant

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/console/status/option"
	"github.com/funtimecoding/soil/pkg/console/status/tag"
)

const (
	DefaultSpaceEnvironment = "CONFLUENCE_DEFAULT_SPACE"
	DefaultPageEnvironment  = "CONFLUENCE_DEFAULT_PAGE"
	LabelEnvironment        = "CONFLUENCE_LABEL"

	NoSpace = "no space"

	PageType = "page"

	Wiki = "/wiki"
	Base = "/wiki/api/v2"

	OldBase = "/wiki/rest/api"
	Search  = "/content/search"
	User    = "/user/current"

	Page     = "/pages"
	Space    = "/spaces"
	Label    = "/labels"
	Children = "/direct-children"

	BodyFormat      = "body-format"
	Status          = "status"
	Query           = "cql"
	SpaceIdentifier = "space-id"
	Title           = "title"

	CurrentStatus = "current"
	Draft         = "draft"
	DraftStatus   = "draft"
	GetDraft      = "get-draft"
)

// Body format
const (
	ViewFormat      = "view"
	AtlasFormat     = "atlas_doc_format"
	StorageFormat   = "storage"
	ExportFormat    = "export_view"
	AnonymousFormat = "anonymous_export_view"
	StyledFormat    = "styled_view"
	EditFormat      = "editor"
)

var (
	ErrorNotFound = errors.New("not found")
	Format        = option.ExtendedColor.Copy()
	Dense         = option.Color.Copy().Tag(tag.Dense)
)
