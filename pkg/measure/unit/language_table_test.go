package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	library "github.com/funtimecoding/soil/pkg/constant"
	module "github.com/funtimecoding/soil/pkg/go_mod/constant"
	"github.com/funtimecoding/soil/pkg/measure/classifier"
	"github.com/funtimecoding/soil/pkg/measure/constant"
	"github.com/funtimecoding/soil/pkg/measure/registry"
	"testing"
)

func TestTableNamesUnique(t *testing.T) {
	seen := map[string]bool{}

	for _, l := range constant.Languages {
		assert.False(t, seen[l.Name])
		seen[l.Name] = true
	}
}

func TestTableExtensionsUnique(t *testing.T) {
	seen := map[string]string{}

	for _, l := range constant.Languages {
		for _, e := range l.Extensions {
			assert.String(t, "", seen[e])
			seen[e] = l.Name
		}

		for _, f := range l.Filenames {
			assert.String(t, "", seen[f])
			seen[f] = l.Name
		}
	}
}

func TestTableCensus(t *testing.T) {
	r := registry.NewDefault()
	assert.String(t, "PHP", r.ByPath("index.php").Name)
	assert.String(t, "Terraform", r.ByPath("main.tf").Name)
	assert.String(t, "Salt", r.ByPath("nginx/init.sls").Name)
	assert.String(t, "YAML", r.ByPath("values.yaml").Name)
	assert.String(t, "YAML", r.ByPath(library.GitLabFile).Name)
	assert.String(t, "Shell", r.ByPath("deploy.sh").Name)
	assert.String(t, "Ruby", r.ByPath("recipes/default.rb").Name)
	assert.String(t, "Ruby", r.ByPath("Berksfile").Name)
	assert.String(t, "ERB", r.ByPath("nginx.conf.erb").Name)
	assert.String(t, "INI", r.ByPath("nginx.conf").Name)
	assert.String(t, "Jinja", r.ByPath("motd.j2").Name)
	assert.String(t, "Dockerfile", r.ByPath(library.ContainerFile).Name)
	assert.String(t, "Systemd", r.ByPath("app.service").Name)
	assert.String(t, "Git Configuration", r.ByPath(".gitignore").Name)
	assert.String(t, "Git Configuration", r.ByPath("chefignore").Name)
	assert.String(t, "Go Template", r.ByPath("deployment.tpl").Name)
	assert.String(t, "XML", r.ByPath("schema.avsc").Name)
	assert.String(t, "Plain Text", r.ByPath("LICENSE").Name)
	assert.String(t, "Groovy", r.ByPath("Jenkinsfile").Name)
	assert.String(t, "Go Module", r.ByPath(module.ModFile).Name)
}

func TestClassifyHashFamily(t *testing.T) {
	c := classifier.New(constant.Markup).Classify(
		`# top
key: value # trailing

list:
  - a
`,
	)
	assert.Integer(t, 1, c.Blank)
	assert.Integer(t, 1, c.Comment)
	assert.Integer(t, 3, c.Code)
}

func TestClassifyDoubleDashFamily(t *testing.T) {
	c := classifier.New(constant.SQL).Classify(
		`-- header
SELECT 1; /* inline */
/*
block
*/
`,
	)
	assert.Integer(t, 4, c.Comment)
	assert.Integer(t, 1, c.Code)
	c = classifier.New(constant.Lua).Classify(
		`--[[ long
comment ]]
print(1) -- trailing
`,
	)
	assert.Integer(t, 2, c.Comment)
	assert.Integer(t, 1, c.Code)
}

func TestClassifySemicolonFamily(t *testing.T) {
	c := classifier.New(constant.INI).Classify(
		`; comment
# also comment
[section]
key = value
`,
	)
	assert.Integer(t, 2, c.Comment)
	assert.Integer(t, 2, c.Code)
}

func TestClassifyHypertextFamily(t *testing.T) {
	c := classifier.New(constant.Hypertext).Classify(
		`<!doctype html>
<!-- one -->
<!--
two
--><p>x</p>
`,
	)
	assert.Integer(t, 3, c.Comment)
	assert.Integer(t, 2, c.Code)
}

func TestClassifyTemplateFamily(t *testing.T) {
	c := classifier.New(constant.Salt).Classify(
		`{# jinja comment #}
# yaml comment
nginx:
  pkg.installed
`,
	)
	assert.Integer(t, 2, c.Comment)
	assert.Integer(t, 2, c.Code)
	c = classifier.New(constant.GoTemplate).Classify(
		`{{/* helm comment */}}
apiVersion: v1
`,
	)
	assert.Integer(t, 1, c.Comment)
	assert.Integer(t, 1, c.Code)
	c = classifier.New(constant.ERB).Classify(
		`<%# erb comment %>
<%= @name %>
`,
	)
	assert.Integer(t, 1, c.Comment)
	assert.Integer(t, 1, c.Code)
}

func TestClassifyPercentFamily(t *testing.T) {
	c := classifier.New(constant.LaTeX).Classify(
		`% preamble
\documentclass{article}
`,
	)
	assert.Integer(t, 1, c.Comment)
	assert.Integer(t, 1, c.Code)
}

func TestClassifyRubyBlock(t *testing.T) {
	c := classifier.New(constant.Ruby).Classify(
		`=begin
doc
=end
puts 1
`,
	)
	assert.Integer(t, 3, c.Comment)
	assert.Integer(t, 1, c.Code)
}

func TestClassifyNoCommentFamily(t *testing.T) {
	c := classifier.New(constant.Notation).Classify(
		`{
  "a": "// not a comment"
}
`,
	)
	assert.Integer(t, 0, c.Comment)
	assert.Integer(t, 3, c.Code)
}
