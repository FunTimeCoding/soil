package constant

import (
	"github.com/funtimecoding/soil/pkg/console/constant"
	"regexp"
)

const (
	ConfluenceDefaultSpaceEnvironment = "CONFLUENCE_DEFAULT_SPACE"
	ConfluenceDefaultPageEnvironment  = "CONFLUENCE_DEFAULT_PAGE"
	ConfluenceLabelEnvironment        = "CONFLUENCE_LABEL"

	ConfluenceNoSpace = "no space"

	ConfluencePageType = "page"

	ConfluenceWiki = "/wiki"
	ConfluenceBase = "/wiki/api/v2"

	ConfluenceOldBase = "/wiki/rest/api"
	ConfluenceSearch  = "/content/search"
	ConfluenceUser    = "/user/current"

	ConfluencePage     = "/pages"
	ConfluenceSpace    = "/spaces"
	ConfluenceLabel    = "/labels"
	ConfluenceChildren = "/direct-children"

	ConfluenceBodyFormat      = "body-format"
	ConfluenceStatus          = "status"
	ConfluenceQuery           = "cql"
	ConfluenceSpaceIdentifier = "space-id"
	ConfluenceTitle           = "title"

	ConfluenceCurrentStatus  = "current"
	ConfluenceDraftParameter = "draft"
	ConfluenceDraftStatus    = "draft"
	ConfluenceGetDraft       = "get-draft"

	ConfluenceExpand = "expand"
	// Body format
	ConfluenceViewFormat      = "view"
	ConfluenceAtlasFormat     = "atlas_doc_format"
	ConfluenceStorageFormat   = "storage"
	ConfluenceExportFormat    = "export_view"
	ConfluenceAnonymousFormat = "anonymous_export_view"
	ConfluenceStyledFormat    = "styled_view"
	ConfluenceEditFormat      = "editor"
)

var (
	ConfluenceFormat = constant.ExtendedColorFormat.Copy()
	ConfluenceDense  = constant.ColorFormat.Copy().Tag(constant.TagDense)
)

var (
	RichTextMacroPattern = regexp.MustCompile(
		`<ac:structured-macro[^>]*ac:name="([^"]+)"[^>]*><ac:rich-text-body>([\s\S]*?)</ac:rich-text-body></ac:structured-macro>`,
	)
	PlainTextMacroPattern = regexp.MustCompile(
		`<ac:structured-macro[^>]*ac:name="([^"]+)"[^>]*><ac:plain-text-body><!\[CDATA\[(.*?)]]></ac:plain-text-body></ac:structured-macro>`,
	)
	MacroCommentPattern = regexp.MustCompile(
		`<!-- ac:(\w+) -->\n([\s\S]*?)<!-- /ac:\w+ -->`,
	)
)
