package face

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/generated/client"
)

type Client interface {
	DeleteGroupName(
		x context.Context,
		name string,
	) (*client.Receipt, error)
	DeleteGroupNameMemberAccount(
		x context.Context,
		name string,
		account string,
	) (*client.Group, error)
	DeleteUserAccount(
		x context.Context,
		account string,
	) (*client.Receipt, error)
	GetGroup(x context.Context) ([]*client.Group, error)
	GetUser(x context.Context) ([]*client.User, error)
	GetUserAccount(
		x context.Context,
		account string,
	) (*client.User, error)
	PatchUserAccount(
		x context.Context,
		account string,
		b client.PatchUserAccountJSONRequestBody,
	) (*client.User, error)
	PostGroup(
		x context.Context,
		b client.PostGroupJSONRequestBody,
	) (*client.Group, error)
	PostGroupNameMember(
		x context.Context,
		name string,
		b client.PostGroupNameMemberJSONRequestBody,
	) (*client.Group, error)
	PostUser(
		x context.Context,
		b client.PostUserJSONRequestBody,
	) (*client.User, error)
	PostUserAccountPassword(
		x context.Context,
		account string,
		b client.PostUserAccountPasswordJSONRequestBody,
	) (*client.Receipt, error)
	Reachable(x context.Context) error
	Close()
}
