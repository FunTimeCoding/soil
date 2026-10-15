package fixture

import (
	"github.com/funtimecoding/soil/pkg/netbox/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"net/http"
	"net/http/httptest"
	"strings"
)

func NewServer(r *Recorder) *httptest.Server {
	return httptest.NewServer(
		http.HandlerFunc(
			func(
				w http.ResponseWriter,
				q *http.Request,
			) {
				w.Header().Set("Content-Type", "application/json")
				path := q.URL.Path

				switch {
				case strings.HasPrefix(path, "/api/dcim/mac-addresses/") &&
					q.Method != http.MethodGet:
					body := decodeBody(q)

					if v, okay := body["assigned_object_type"].(string); okay {
						r.Assigned = v
					}

					respond(w, constant.FixtureMacRecord)
				case path == "/api/dcim/mac-addresses/":
					respond(w, constant.FixtureMacList)
				case strings.HasPrefix(
					path,
					"/api/virtualization/interfaces/",
				) &&
					q.Method == http.MethodPut:
					body := decodeBody(q)

					if v, okay := body["primary_mac_address"]; okay && v != nil {
						r.Primary = "set"
					}

					respond(w, constant.FixtureVirtualInterface)
				case path == "/api/virtualization/interfaces/":
					respond(w, constant.FixtureVirtualInterfaceList)
				case path == "/api/extras/config-contexts/1/" &&
					q.Method == http.MethodPatch:
					r.Patched = decodeBody(q)
					respond(w, constant.FixtureConfigurationContext)
				case path == "/api/extras/config-contexts/":
					respond(w, constant.FixtureConfigurationContextList)
				case path == "/api/extras/custom-field-choice-sets/" &&
					q.Method == http.MethodPost:
					r.Created = decodeBody(q)
					respond(w, constant.FixtureChoiceSet)
				case path == "/api/extras/custom-field-choice-sets/":
					respond(w, constant.FixtureChoiceSetList)
				case path == "/api/extras/custom-fields/" &&
					q.Method == http.MethodPost:
					r.Created = decodeBody(q)
					respond(w, constant.FixtureCustomField)
				default:
					r.Unmatched = append(
						r.Unmatched,
						join.Space(q.Method, path),
					)
					w.WriteHeader(http.StatusNotFound)
				}
			},
		),
	)
}
