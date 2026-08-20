package common

import (
	"github.com/funtimecoding/soil/pkg/atlassian/constant"
	"github.com/funtimecoding/soil/pkg/atlassian/opsgenie"
	"github.com/funtimecoding/soil/pkg/atlassian/opsgenie/alert/alert_enricher"
	"github.com/funtimecoding/soil/pkg/atlassian/opsgenie/alert/detail"
	prometheus "github.com/funtimecoding/soil/pkg/prometheus/constant"
	stringConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/strings/split"
	"github.com/funtimecoding/soil/pkg/strings/split/key_value"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"strings"
)

func Opsgenie() *opsgenie.Client {
	result := opsgenie.NewEnvironment().Set(
		alert_enricher.New().Add(
			prometheus.HighMemoryUsage,
			prometheus.EntityMemory,
			prometheus.CategoryHigh,
		),
	)

	if s := environment.Optional(constant.OpsgenieTeamEnvironment); s != "" {
		for _, pair := range split.Semicolon(s) {
			if len(pair) == 0 {
				continue
			}

			k, v := key_value.Comma(pair)

			if len(k) == 0 || len(v) == 0 {
				continue
			}

			result.TeamMap().AddKey(k, v)
		}
	} else {
		result.TeamMap().AddKey("Infinite Loopsies", "INF")
	}

	result.ShortAlert(
		func(s string) string {
			if false {
				switch s {
				case prometheus.HighMemoryUsage:
					return "Memory"
				}
			}

			return s
		},
	)
	result.ShortUser(
		func(s string) string {
			if strings.Contains(s, stringConstant.At) {
				k, _ := key_value.At(s)

				return k
			}

			return s
		},
	)
	result.DescriptionToName(
		func(s string) string {
			return s
		},
	)
	result.TagToTeam(
		func(s []string) string {
			return ""
		},
	)
	result.ParseDescription(
		func(s string) *detail.Prometheus {
			return detail.New()
		},
	)

	return result
}
