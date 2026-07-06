package scan

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/parse"
)

func checkIdentitySource(
	content string,
	file string,
	path string,
	name string,
) []*concern.Concern {
	f, _, e := parse.Source(file, content)

	if e != nil {
		return []*concern.Concern{
			concern.NewPackage(
				IdentityNotParseableKey,
				IdentityNotParseableText,
				path,
			),
		}
	}

	v := parse.FindVariable(f, "Identity")

	if v == nil {
		return []*concern.Concern{
			concern.NewPackage(
				IdentityMissingVariableKey,
				IdentityMissingVariableText,
				path,
			),
		}
	}

	if len(v.Values) == 0 {
		return []*concern.Concern{
			concern.NewPackage(
				IdentityNoInitializerKey,
				IdentityNoInitializerText,
				path,
			),
		}
	}

	c := parse.FindCall(v.Values[0], "identity", "New")

	if c == nil {
		return []*concern.Concern{
			concern.NewPackage(
				IdentityNotNewKey,
				IdentityNotNewText,
				path,
			),
		}
	}

	if len(c.Args) < 1 {
		return []*concern.Concern{
			concern.NewPackage(
				IdentityNoArgumentKey,
				IdentityNoArgumentText,
				path,
			),
		}
	}

	declared, okay := parse.StringValue(c.Args[0])

	if !okay {
		return []*concern.Concern{
			concern.NewPackage(
				IdentityNotStringKey,
				IdentityNotStringText,
				path,
			),
		}
	}

	if declared != name {
		return []*concern.Concern{
			concern.NewPackage(
				IdentityMismatchKey,
				fmt.Sprintf(
					"Identity name %s does not match directory %s",
					declared,
					name,
				),
				path,
			),
		}
	}

	return nil
}
