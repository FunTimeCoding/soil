package service

import "github.com/funtimecoding/soil/pkg/tool/gopostgresd/constant"

func schemaOrDefault(schema string) string {
	if schema == "" {
		return constant.DefaultSchema
	}

	return schema
}
