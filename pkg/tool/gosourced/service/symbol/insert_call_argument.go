package symbol

import "github.com/dave/dst"

func insertCallArgument(
	c *dst.CallExpr,
	position int,
	a dst.Expr,
) {
	multiline := false
	var arguments []dst.Expr

	for i, existing := range c.Args {
		if existing.Decorations().Before == dst.NewLine {
			multiline = true
		}

		if i == position {
			arguments = append(arguments, a)
		}

		arguments = append(arguments, existing)
	}

	if position >= len(c.Args) {
		arguments = append(arguments, a)
	}

	if multiline {
		for i, argument := range arguments {
			d := argument.Decorations()
			d.Before = dst.NewLine
			d.After = dst.None

			if i == len(arguments)-1 {
				d.After = dst.NewLine
			}
		}
	}

	c.Args = arguments
}
