package unit

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gosublimed/face"
	"github.com/funtimecoding/soil/pkg/tool/gosublimed/generated/client"
)

func viewAnswers(
	x context.Context,
	f face.Client,
) []string {
	var result []string
	answer := func(
		v *client.View,
		e error,
	) {
		result = append(result, answerOf(v, e))
	}
	answer(f.GetView(x, 1))
	answer(
		f.CreateView(
			x,
			client.CreateViewJSONRequestBody{
				Title:   "alfa",
				Content: "bravo charlie bravo",
				Syntax:  new("text"),
			},
		),
	)
	answer(f.GetView(x, 1))
	answer(
		f.EditView(
			x,
			1,
			client.EditViewJSONRequestBody{
				OldString: "bravo",
				NewString: "delta",
			},
		),
	)
	answer(f.EditView(x, 1, client.EditViewJSONRequestBody{NewString: "delta"}))
	answer(
		f.EditView(
			x,
			1,
			client.EditViewJSONRequestBody{
				OldString:  "bravo",
				NewString:  "delta",
				ReplaceAll: new(true),
			},
		),
	)
	answer(
		f.EditView(
			x,
			9,
			client.EditViewJSONRequestBody{
				OldString: "bravo",
				NewString: "delta",
			},
		),
	)
	answer(
		nil,
		f.SaveView(
			x,
			1,
			client.SaveViewJSONRequestBody{FilePath: new("/alfa/bravo.txt")},
		),
	)
	answer(f.GetView(x, 1))
	answer(nil, f.SaveView(x, 9, client.SaveViewJSONRequestBody{}))
	answer(
		f.OpenFile(x, client.OpenFileJSONRequestBody{FilePath: "/charlie.txt"}),
	)
	views, e := f.GetViews(x)
	answer(nil, e)

	for _, v := range views {
		answer(&v, nil)
	}

	answer(nil, f.CloseView(x, 1))
	answer(f.GetView(x, 1))
	answer(nil, f.CloseView(x, 1))
	views, e = f.GetViews(x)
	answer(nil, e)

	for _, v := range views {
		answer(&v, nil)
	}

	return result
}
