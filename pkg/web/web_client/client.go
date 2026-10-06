package web_client

import (
	"github.com/funtimecoding/soil/pkg/face"
	"net/http"
)

type Client struct {
	clock  face.Clock
	client *http.Client
}
