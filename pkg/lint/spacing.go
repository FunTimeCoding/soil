package lint

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"github.com/funtimecoding/soil/pkg/lint/file_report"
	"io"
	"strings"
)

func Spacing(
	path string,
	r io.Reader,
) *file_report.Report {
	s := file_report.New(path, r)
	var pastLine string
	var pastWasBlank bool
	var needBlankAfterClosingBrace bool
	var inBacktick bool
	var blockStack []bool // true = control block
	var pendingControl bool
	var parenDepth int
	var pendingBlank bool
	var pendingBlankLine int

	for s.Scan() {
		line, number := s.Text()
		trimmed := strings.TrimSpace(line)
		isBlank := trimmed == ""
		wasInBacktick := inBacktick
		topLevel := len(blockStack) == 0

		if strings.Count(line, "`")%2 == 1 {
			inBacktick = !inBacktick
		}

		if wasInBacktick {
			s.ChangedLine(line)
			pastLine = line
			pastWasBlank = isBlank

			continue
		}

		isControlStart := strings.HasPrefix(trimmed, "if ") ||
			strings.HasPrefix(trimmed, "for ") ||
			strings.HasPrefix(trimmed, "switch ") ||
			strings.HasPrefix(trimmed, "select ")

		isExit := trimmed == "return" || strings.HasPrefix(
			trimmed,
			"return ",
		) ||
			trimmed == "break" || strings.HasPrefix(trimmed, "break ") ||
			trimmed == "continue" || strings.HasPrefix(trimmed, "continue ")

		isDefer := strings.HasPrefix(trimmed, "defer ")

		isTopLevelDeclaration := topLevel && !isBlank && (strings.HasPrefix(
			trimmed,
			"func ",
		) ||
			strings.HasPrefix(trimmed, "type "))

		isVariable := topLevel && !isBlank && strings.HasPrefix(
			trimmed,
			"var ",
		)
		isConstant := topLevel &&
			!isBlank &&
			strings.HasPrefix(trimmed, "const ")

		isClosingBrace := strings.HasPrefix(trimmed, "}") ||
			strings.HasPrefix(trimmed, "case ") ||
			trimmed == "default:"

		pastTrimmed := strings.TrimSpace(pastLine)
		pastOpensBlock := strings.HasSuffix(pastTrimmed, "{") ||
			strings.HasPrefix(pastTrimmed, "case ") ||
			pastTrimmed == "default:" ||
			strings.HasPrefix(pastTrimmed, "//")

		isElseContinuation := strings.HasPrefix(trimmed, "} else")
		endsWithBrace := strings.HasSuffix(trimmed, "{")

		if isElseContinuation {
			if len(blockStack) > 0 {
				blockStack = blockStack[:len(blockStack)-1]
			}

			if endsWithBrace {
				blockStack = append(blockStack, true)
			}

			pendingControl = false
		} else if strings.HasPrefix(trimmed, "}") && endsWithBrace {
			if len(blockStack) > 0 {
				blockStack = blockStack[:len(blockStack)-1]
			}

			blockStack = append(blockStack, false)
			pendingControl = false
		} else if endsWithBrace {
			blockStack = append(blockStack, isControlStart || pendingControl)
			pendingControl = false
		} else if isControlStart {
			pendingControl = true
		} else if !isBlank && !isClosingBrace && parenDepth == 0 {
			pastContinues := strings.HasSuffix(pastTrimmed, "&&") ||
				strings.HasSuffix(pastTrimmed, "||")

			if !pendingControl || !pastContinues {
				pendingControl = false
			}
		}

		for _, c := range line {
			switch c {
			case '(':
				parenDepth++
			case ')':
				if parenDepth > 0 {
					parenDepth--
				}
			}
		}

		// Decide on a held blank when the next non-blank line arrives.
		if !isBlank && pendingBlank {
			if strings.HasPrefix(trimmed, "//") {
				// Blank before a comment is always valid - emit both immediately.
				// Clear needBlankAfterClosingBrace: the emitted blank satisfies it.
				s.ChangedLine("")
				s.ChangedLine(line)
				pendingBlank = false
				needBlankAfterClosingBrace = false

				continue
			}

			pendingBlank = false

			if topLevel {
				pastIsVariable := strings.HasPrefix(pastTrimmed, "var ")
				pastIsConstant := strings.HasPrefix(pastTrimmed, "const ")
				sameKind := (pastIsVariable && isVariable) || (pastIsConstant && isConstant)

				if sameKind {
					// Spurious blank between consecutive top-level declarations of the same kind - remove.
					s.AddConcern(
						constant.ExtraneousTopLevelBlankKey,
						constant.ExtraneousTopLevelBlankText,
						path,
						pendingBlankLine,
						"",
						true,
					)
				} else {
					s.ChangedLine("")
				}
			} else {
				// Blank is invalid after a block opener or a comment - a
				// comment is attached to what follows; the blank belongs before it.
				pastOpenedBrace := strings.HasSuffix(pastTrimmed, "{") ||
					strings.HasPrefix(pastTrimmed, "case ") ||
					pastTrimmed == "default:" ||
					strings.HasPrefix(pastTrimmed, "//")

				if pastOpenedBrace {
					// Blank at start of block - always invalid.
					s.AddConcern(
						constant.BlankInsideFunctionKey,
						constant.BlankInsideFunctionText,
						path,
						pendingBlankLine,
						"",
						true,
					)
				} else if needBlankAfterClosingBrace {
					// Blank earned by a preceding control block - emit it and
					// clear the flag so the existing insertion check won't fire.
					s.ChangedLine("")
					needBlankAfterClosingBrace = false
				} else if isControlStart || isExit || isDefer {
					// Blank before control, exit, or defer - valid, emit.
					s.ChangedLine("")
				} else {
					// Any other blank inside a function body - invalid.
					s.AddConcern(
						constant.BlankInsideFunctionKey,
						constant.BlankInsideFunctionText,
						path,
						pendingBlankLine,
						"",
						true,
					)
				}
			}
		}

		if isControlStart && !pastWasBlank && pastLine != "" && !pastOpensBlock {
			s.ChangedLine("")
			s.AddConcern(
				constant.MissingBlankBeforeControlKey,
				constant.MissingBlankBeforeControlText,
				path,
				number,
				line,
				true,
			)
			needBlankAfterClosingBrace = false
		}

		if isExit && !pastWasBlank && pastLine != "" && !pastOpensBlock {
			s.ChangedLine("")
			s.AddConcern(
				constant.MissingBlankBeforeExitKey,
				constant.MissingBlankBeforeExitText,
				path,
				number,
				line,
				true,
			)
			needBlankAfterClosingBrace = false
		}

		if isTopLevelDeclaration && !pastWasBlank && pastLine != "" && !pastOpensBlock {
			s.ChangedLine("")
			s.AddConcern(
				constant.MissingBlankBeforeDeclarationKey,
				constant.MissingBlankBeforeDeclarationText,
				path,
				number,
				line,
				true,
			)
		}

		pastIsVariable := strings.HasPrefix(pastTrimmed, "var ")
		pastIsConstant := strings.HasPrefix(pastTrimmed, "const ")
		crossKind := (pastIsVariable && isConstant) || (pastIsConstant && isVariable)

		if crossKind && !pastWasBlank && pastLine != "" {
			s.ChangedLine("")
			s.AddConcern(
				constant.MissingBlankBetweenVariableConstantKey,
				constant.MissingBlankBetweenVariableConstantText,
				path,
				number,
				line,
				true,
			)
		}

		if needBlankAfterClosingBrace && !isBlank && !isClosingBrace {
			s.ChangedLine("")
			s.AddConcern(
				constant.MissingBlankAfterControlKey,
				constant.MissingBlankAfterControlText,
				path,
				number,
				line,
				true,
			)
		}

		if isBlank && pastWasBlank {
			s.AddConcern(
				constant.ExtraneousBlankLineKey,
				constant.ExtraneousBlankLineText,
				path,
				number,
				line,
				true,
			)

			if !pendingBlank {
				needBlankAfterClosingBrace = false
			}

			continue
		}

		// Hold the blank to defer the decision to the next non-blank line.
		if isBlank {
			pendingBlank = true
			pendingBlankLine = number
			pastWasBlank = true

			// Do not update pastLine - keep it as the last non-blank line so
			// pastOpensBlock is correct when the next non-blank is processed.
			continue
		}

		s.ChangedLine(line)

		if strings.HasPrefix(
			trimmed,
			"}",
		) && !isElseContinuation && !endsWithBrace {
			isControl := false

			if len(blockStack) > 0 {
				isControl = blockStack[len(blockStack)-1]
				blockStack = blockStack[:len(blockStack)-1]
			}

			if trimmed == "}" {
				needBlankAfterClosingBrace = isControl
			} else {
				needBlankAfterClosingBrace = false
			}
		} else {
			needBlankAfterClosingBrace = false
		}

		pastLine = line
		pastWasBlank = isBlank
	}

	return s.Finalize()
}
