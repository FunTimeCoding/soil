package constant

import (
	"github.com/funtimecoding/soil/pkg/lint/types/suggestion"
	"github.com/funtimecoding/soil/pkg/maps"
)

var Suggestions = maps.Join(CodeSuggestions, FileSuggestions, SystemSuggestions)
var CodeSuggestions = map[string]suggestion.Suggestion{
	"tx":     {Letters: []string{"t"}},
	"ctx":    {Letters: []string{"x"}},
	"param":  {Letters: []string{"p"}, Words: []string{"parameter"}},
	"pos":    {Letters: []string{"p"}, Words: []string{"position"}},
	"buf":    {Letters: []string{"b"}, Words: []string{"buffer"}},
	"ptr":    {Letters: []string{"p"}, Words: []string{"pointer"}},
	"ref":    {Letters: []string{"r"}, Words: []string{"reference"}},
	"prev":   {Words: []string{"past"}},
	"decl":   {Words: []string{"declaration"}},
	"elem":   {Letters: []string{"e"}, Words: []string{"element"}},
	"def":    {Words: []string{"definition"}},
	"concat": {Words: []string{"concatenate"}},
	"obj":    {Letters: []string{"o"}, Words: []string{"object"}},
	"stmt":   {Letters: []string{"s"}, Words: []string{"statement"}},
	"var":    {Letters: []string{"v"}, Words: []string{"variable"}},
	"const":  {Letters: []string{"c"}, Words: []string{"constant"}},
	"func":   {Letters: []string{"f"}, Words: []string{"function"}},
	"fn":     {Letters: []string{"f"}, Words: []string{"function"}},
	"sig":    {Letters: []string{"s"}, Words: []string{"signature"}},
	"pkg":    {Letters: []string{"p"}, Words: []string{"package"}},
	"err":    {Letters: []string{"e", "f", "g", "h"}, Words: []string{"error"}},
	"ok":     {Words: []string{"okay"}},
	"val":    {Letters: []string{"v"}, Words: []string{"value"}},
	"args":   {Words: []string{"arguments"}},
	"opt":    {Letters: []string{"o"}, Words: []string{"option"}},
	"cols":   {Words: []string{"columns"}},
	"dep":    {Letters: []string{"d"}, Words: []string{"dependency"}},
}

var FileSuggestions = map[string]suggestion.Suggestion{
	"dir":    {Letters: []string{"d"}, Words: []string{"directory"}},
	"dirs":   {Words: []string{"directories"}},
	"src":    {Letters: []string{"s"}, Words: []string{"source"}},
	"dst":    {Letters: []string{"d"}, Words: []string{"destination"}},
	"doc":    {Letters: []string{"d"}, Words: []string{"document"}},
	"yaml":   {Letters: []string{"m"}, Words: []string{"markup"}},
	"xml":    {Letters: []string{"m"}, Words: []string{"markup"}},
	"html":   {Letters: []string{"m"}, Words: []string{"markup"}},
	"json":   {Letters: []string{"j"}, Words: []string{"notation"}},
	"config": {Letters: []string{"c"}, Words: []string{"configuration"}},
	"cfg":    {Letters: []string{"c"}, Words: []string{"configuration"}},
	"tmp":    {Letters: []string{"t"}},
}

var SystemSuggestions = map[string]suggestion.Suggestion{
	"id":   {Words: []string{"identifier"}},
	"url":  {Letters: []string{"l"}, Words: []string{"locator"}},
	"mcp":  {Letters: []string{"c"}, Words: []string{"model_context"}},
	"msg":  {Letters: []string{"m"}, Words: []string{"message"}},
	"req":  {Letters: []string{"r"}, Words: []string{"request"}},
	"addr": {Letters: []string{"a"}, Words: []string{"address"}},
	"nav":  {Letters: []string{"n"}, Words: []string{"navigation"}},
	"llm":  {Letters: []string{"m"}, Words: []string{"model"}},
	"href": {Words: []string{"link", "reference", "locator"}},
	"loc":  {Letters: []string{"l"}, Words: []string{"location"}},
	"os":   {Words: []string{"operating_system", "system"}},
	"osd":  {Words: []string{"storage_daemon"}},
	"fqdn": {Words: []string{"absolute_domain_name", "absolute_name"}},
}
