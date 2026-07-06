package scan

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/system/virtual_file_system"
	"gopkg.in/yaml.v3"
	"path/filepath"
	"sort"
	"strings"
)

func errorHandlingConcerns(
	v *virtual_file_system.System,
	path string,
	serviceName string,
	suppress map[string][]string,
) []*concern.Concern {
	file := filepath.Join(path, "generated", "server", "openapi.yaml")

	if !v.Has(file) {
		return nil
	}

	var spec errorHandlingSpec

	if yaml.Unmarshal(v.Read(file), &spec) != nil {
		return nil
	}

	var result []*concern.Concern
	referencesError := false
	referencesErrorResponse := false

	for _, methods := range spec.Paths {
		for _, op := range methods {
			for code, response := range op.Responses {
				if strings.HasPrefix(code, "2") {
					continue
				}

				if response.Reference != "" {
					result = append(
						result,
						concern.NewPackage(
							ResponseReferenceKey,
							fmt.Sprintf(
								"%s %s: uses components/responses ref (should be inlined)",
								op.OperationID,
								code,
							),
							path,
						),
					)

					continue
				}

				r := schemaReference(response)

				if r != "" {
					if r == "Error" {
						referencesError = true
					}

					if r == "ErrorResponse" {
						referencesErrorResponse = true
					}
				}

				if r == "" {
					result = append(
						result,
						concern.NewPackage(
							BodylessErrorKey,
							fmt.Sprintf(
								"%s %s: bodyless %s response",
								op.OperationID,
								code,
								response.Description,
							),
							path,
						),
					)

					continue
				}

				if code == "500" && r == "Error" {
					result = append(
						result,
						concern.NewPackage(
							ErrorOnServerErrorKey,
							fmt.Sprintf(
								"%s: 500 references Error (expected ErrorResponse)",
								op.OperationID,
							),
							path,
						),
					)
				}

				if code != "500" && r == "ErrorResponse" {
					if isSuppressed(
						suppress,
						op.OperationID,
						ErrorResponseOnNon500Key,
					) {
						continue
					}

					result = append(
						result,
						concern.NewPackage(
							ErrorResponseOnNon500Key,
							fmt.Sprintf(
								"%s: %s references ErrorResponse (expected Error)",
								op.OperationID,
								code,
							),
							path,
						),
					)
				}
			}
		}
	}

	if referencesError {
		if _, okay := spec.Components.Schemas["Error"]; !okay {
			result = append(
				result,
				concern.NewPackage(
					MissingErrorSchemaKey,
					"endpoints reference Error but schema not defined",
					path,
				),
			)
		}
	}

	if referencesErrorResponse {
		if _, okay := spec.Components.Schemas["ErrorResponse"]; !okay {
			result = append(
				result,
				concern.NewPackage(
					MissingErrorResponseSchemaKey,
					"endpoints reference ErrorResponse but schema not defined",
					path,
				),
			)
		}
	}

	sort.Slice(
		result,
		func(i, j int) bool {
			return result[i].Text < result[j].Text
		},
	)

	return result
}
