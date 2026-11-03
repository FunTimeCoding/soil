package unit

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/face"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/generated/client"
)

func directoryAnswers(
	x context.Context,
	f face.Client,
) []string {
	var result []string
	answer := func(
		v any,
		e error,
	) {
		result = append(result, answerOf(v, e))
	}
	alfa := client.PostUserJSONRequestBody{
		Account:  "alfa",
		Name:     "Alfa",
		Surname:  "Bravo",
		Mail:     new("alfa@example.org"),
		Password: new("charlie"),
	}
	member := client.PostGroupNameMemberJSONRequestBody{Account: "alfa"}
	answer(f.GetUser(x))
	answer(f.GetUserAccount(x, "alfa"))
	answer(f.PostUser(x, alfa))
	answer(f.PostUser(x, alfa))
	answer(f.GetUserAccount(x, "alfa"))
	answer(
		f.PatchUserAccount(
			x,
			"alfa",
			client.PatchUserAccountJSONRequestBody{Name: new("Delta")},
		),
	)
	answer(
		f.PatchUserAccount(
			x,
			"missing",
			client.PatchUserAccountJSONRequestBody{},
		),
	)
	answer(
		f.PostUserAccountPassword(
			x,
			"alfa",
			client.PostUserAccountPasswordJSONRequestBody{Password: "echo"},
		),
	)
	answer(
		f.PostUserAccountPassword(
			x,
			"missing",
			client.PostUserAccountPasswordJSONRequestBody{Password: "echo"},
		),
	)
	answer(f.PostGroup(x, client.PostGroupJSONRequestBody{Name: "echo"}))
	answer(f.PostGroup(x, client.PostGroupJSONRequestBody{Name: "echo"}))
	answer(f.PostGroup(x, client.PostGroupJSONRequestBody{Name: "foxtrot"}))
	answer(f.PostGroupNameMember(x, "echo", member))
	answer(f.PostGroupNameMember(x, "echo", member))
	answer(
		f.PostGroupNameMember(
			x,
			"echo",
			client.PostGroupNameMemberJSONRequestBody{Account: "missing"},
		),
	)
	answer(f.PostGroupNameMember(x, "missing", member))
	answer(f.GetGroup(x))
	answer(f.DeleteGroupNameMemberAccount(x, "echo", "alfa"))
	answer(f.DeleteGroupNameMemberAccount(x, "echo", "alfa"))
	answer(f.DeleteGroupName(x, "foxtrot"))
	answer(f.DeleteGroupName(x, "foxtrot"))
	answer(f.DeleteUserAccount(x, "alfa"))
	answer(f.DeleteUserAccount(x, "alfa"))
	answer(f.GetUser(x))
	answer(f.GetGroup(x))

	return result
}
