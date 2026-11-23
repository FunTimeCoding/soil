package graph

import (
	"github.com/funtimecoding/soil/pkg/microsoft/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"golang.org/x/oauth2"
	"strings"
)

func refusalCode(r *oauth2.RetrieveError) string {
	return strings.TrimSpace(
		join.Space(
			r.ErrorCode,
			constant.DirectoryCode.FindString(r.ErrorDescription),
		),
	)
}
