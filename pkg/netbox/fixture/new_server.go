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
				case path == "/api/dcim/devices/3/" &&
					q.Method == http.MethodPatch:
					r.Patched = decodeBody(q)
					respond(w, constant.FixtureDevice)
				case path == "/api/dcim/devices/":
					respond(w, constant.FixtureDeviceList)
				case path == "/api/dcim/device-types/":
					respond(w, constant.FixtureDeviceTypeList)
				case path == "/api/dcim/device-roles/":
					respond(w, constant.FixtureDeviceRoleList)
				case path == "/api/dcim/sites/":
					respond(w, constant.FixtureSiteList)
				case path == "/api/dcim/power-outlets/" &&
					q.Method == http.MethodPost:
					r.Created = decodeBody(q)
					respond(w, constant.FixturePowerOutlet)
				case path == "/api/dcim/power-outlets/":
					respond(w, constant.FixturePowerOutletList)
				case path == "/api/dcim/power-ports/" &&
					q.Method == http.MethodPost:
					r.Created = decodeBody(q)
					respond(w, constant.FixturePowerPort)
				case path == "/api/dcim/interfaces/8/" &&
					q.Method == http.MethodPatch:
					r.Patched = decodeBody(q)
					respond(w, constant.FixtureInterface)
				case path == "/api/dcim/interfaces/":
					respond(w, constant.FixtureInterfaceList)
				case strings.HasPrefix(path, "/api/dcim/cables/") &&
					q.Method == http.MethodDelete:
					r.Deleted = append(r.Deleted, path)
					w.WriteHeader(http.StatusNoContent)
				case path == "/api/dcim/cables/" && q.Method == http.MethodPost:
					r.Created = decodeBody(q)
					respond(w, constant.FixtureCable)
				case path == "/api/virtualization/virtual-machines/2/" &&
					q.Method == http.MethodPatch:
					r.Patched = decodeBody(q)
					respond(w, constant.FixtureVirtualMachine)
				case path == "/api/virtualization/virtual-machines/":
					respond(w, constant.FixtureVirtualMachineList)
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
