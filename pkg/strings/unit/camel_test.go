package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/strings/camel"
	"testing"
)

func TestCamelToSnake(t *testing.T) {
	assert.String(t, "add_file_to_tar", camel.ToSnake("addFileToTar"))
	assert.String(t, "serve_http", camel.ToSnake("ServeHTTP"))
	assert.String(t, "http_server", camel.ToSnake("HTTPServer"))
	assert.String(t, "name_2", camel.ToSnake("name2"))
	assert.String(t, "already_snake", camel.ToSnake("already_snake"))
	assert.String(t, "new", camel.ToSnake("New"))
}
