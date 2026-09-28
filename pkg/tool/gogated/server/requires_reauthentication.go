package server

import (
	"github.com/funtimecoding/soil/pkg/tool/gogated/constant"
	"net/http"
	"slices"
	"strconv"
	"strings"
	"time"
)

func requiresReauthentication(
	r *http.Request,
	authenticatedAt time.Time,
) bool {
	query := r.URL.Query()

	if slices.Contains(
		strings.Fields(query.Get(constant.PromptParameter)),
		constant.PromptLogin,
	) {
		return true
	}

	raw := query.Get(constant.MaxAgeParameter)

	if raw == "" {
		return false
	}

	maxAge, e := strconv.ParseInt(raw, 10, 64)

	if e != nil {
		return false
	}

	if maxAge <= 0 {
		return true
	}

	return time.Since(authenticatedAt) >
		time.Duration(maxAge)*time.Second
}
