package client

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/google/uuid"
)

func generateVerifier() string {
	return join.Empty(uuid.New().String(), uuid.New().String())
}
