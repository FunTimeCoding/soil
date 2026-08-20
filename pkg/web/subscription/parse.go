package subscription

import (
	"github.com/funtimecoding/soil/pkg/web/constant"
	"net/http"
	"strings"
)

func Parse(r *http.Request) Subscription {
	value := r.URL.Query().Get(constant.ParameterSubscribe)

	if value == "" {
		return Subscription{}
	}

	result := Subscription{}

	for _, s := range strings.Split(value, ",") {
		result[s] = true
	}

	return result
}
