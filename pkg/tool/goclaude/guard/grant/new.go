package grant

import "github.com/funtimecoding/soil/pkg/tool/goclauded/face"

func New(
	client face.Client,
	session string,
	directory string,
) *Grant {
	return &Grant{client: client, session: session, directory: directory}
}
