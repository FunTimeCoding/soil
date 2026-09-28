package server

import (
	"github.com/funtimecoding/soil/pkg/tool/gogated/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"net/http"
)

func (s *Server) renderLoginForm(
	w http.ResponseWriter,
	errorMessage string,
	source string,
) {
	var content []gomponents.Node

	if errorMessage != "" {
		content = append(
			content,
			html.P(
				html.Style("color: var(--pico-del-color);"),
				gomponents.Text(errorMessage),
			),
		)
	}

	directory := s.service.DirectoryConfigured()
	mailType := "email"
	mailLabel := constant.MailLabel

	if directory {
		mailType = "text"
		mailLabel = constant.AccountLabel
	}

	form := []gomponents.Node{
		html.Method("post"),
		html.Action("/authorize"),
		html.Label(html.For(constant.MailField), gomponents.Text(mailLabel)),
		html.Input(
			html.Type(mailType),
			html.ID(constant.MailField),
			html.Name(constant.MailField),
			gomponents.Attr("required", ""),
			gomponents.Attr("autofocus", ""),
		),
		html.Label(
			html.For(constant.PasswordField),
			gomponents.Text("Password"),
		),
		html.Input(
			html.Type(constant.PasswordInputType),
			html.ID(constant.PasswordField),
			html.Name(constant.PasswordField),
			gomponents.Attr("required", ""),
		),
	}

	if directory {
		form = append(
			form,
			sourceChoice(
				constant.SourceDirectory,
				constant.DirectoryLabel,
				source,
			),
			sourceChoice(constant.SourceLocal, constant.LocalLabel, source),
		)
	}

	form = append(
		form,
		html.Button(html.Type("submit"), gomponents.Text("Sign in")),
	)
	content = append(content, html.Form(form...))
	s.view.RenderPage(w, "Sign in", "/authorize", content...)
}
