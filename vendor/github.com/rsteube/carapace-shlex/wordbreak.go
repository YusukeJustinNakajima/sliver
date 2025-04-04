package shlex

import "encoding/json"

const BASH_WORDBREAKS = " \t\r\n" + `"'><=;|&(:`

type WordbreakType int

const (
	WORDBREAK_UNKNOWN WordbreakType = iota
	// https://www.gnu.org/software/bash/manual/html_node/Redirections.html
	WORDBREAK_REDIRECT_INPUT
	WORDBREAK_REDIRECT_OUTPUT
	WORDBREAK_REDIRECT_OUTPUT_APPEND
	WORDBREAK_REDIRECT_OUTPUT_BOTH
	WORDBREAK_REDIRECT_OUTPUT_BOTH_APPEND
	WORDBREAK_REDIRECT_INPUT_STRING
	WORDBREAK_REDIRECT_INPUT_DUPLICATE
	WORDBREAK_REDIRECT_INPUT_OUTPUT
	// https://www.gnu.org/software/bash/manual/html_node/Pipelines.html
	WORDBREAK_PIPE
	WORDBREAK_PIPE_WITH_STDERR
	// https://www.gnu.org/software/bash/manual/html_node/Lists.html)
	WORDBREAK_LIST_ASYNC
	WORDBREAK_LIST_SEQUENTIAL
	WORDBREAK_LIST_AND
	WORDBREAK_LIST_OR
	// COMP_WORDBREAKS
	WORDBREAK_CUSTOM
)

var wordbreakTypes = map[WordbreakType]string{
	WORDBREAK_UNKNOWN:                     "WORDBREAK_UNKNOWN",
	WORDBREAK_REDIRECT_INPUT:              "WORDBREAK_REDIRECT_INPUT",
	WORDBREAK_REDIRECT_OUTPUT:             "WORDBREAK_REDIRECT_OUTPUT",
	WORDBREAK_REDIRECT_OUTPUT_APPEND:      "WORDBREAK_REDIRECT_OUTPUT_APPEND",
	WORDBREAK_REDIRECT_OUTPUT_BOTH:        "WORDBREAK_REDIRECT_OUTPUT_BOTH",
	WORDBREAK_REDIRECT_OUTPUT_BOTH_APPEND: "WORDBREAK_REDIRECT_OUTPUT_BOTH_APPEND",
	WORDBREAK_REDIRECT_INPUT_STRING:       "WORDBREAK_REDIRECT_INPUT_STRING",
	WORDBREAK_REDIRECT_INPUT_DUPLICATE:    "WORDBREAK_REDIRECT_INPUT_DUPLICATE",
	WORDBREAK_REDIRECT_INPUT_OUTPUT:       "WORDBREAK_REDIRECT_INPUT_OUTPUT",
	WORDBREAK_PIPE:                        "WORDBREAK_PIPE",
	WORDBREAK_PIPE_WITH_STDERR:            "WORDBREAK_PIPE_WITH_STDERR",
	WORDBREAK_LIST_ASYNC:                  "WORDBREAK_LIST_ASYNC",
	WORDBREAK_LIST_SEQUENTIAL:             "WORDBREAK_LIST_SEQUENTIAL",
	WORDBREAK_LIST_AND:                    "WORDBREAK_LIST_AND",
	WORDBREAK_LIST_OR:                     "WORDBREAK_LIST_OR",
	WORDBREAK_CUSTOM:                      "WORDBREAK_CUSTOM",
}

func (w WordbreakType) MarshalJSON() ([]byte, error) {
	return json.Marshal(wordbreakTypes[w])
}

func (w WordbreakType) IsPipelineDelimiter() bool {
	switch w {
	case
		WORDBREAK_PIPE,
		WORDBREAK_PIPE_WITH_STDERR,
		WORDBREAK_LIST_ASYNC,
		WORDBREAK_LIST_SEQUENTIAL,
		WORDBREAK_LIST_AND,
		WORDBREAK_LIST_OR:
		return true
	default:
		return false
	}
}

func (w WordbreakType) IsRedirect() bool {
	switch w {
	case
		WORDBREAK_REDIRECT_INPUT,
		WORDBREAK_REDIRECT_OUTPUT,
		WORDBREAK_REDIRECT_OUTPUT_APPEND,
		WORDBREAK_REDIRECT_OUTPUT_BOTH,
		WORDBREAK_REDIRECT_OUTPUT_BOTH_APPEND,
		WORDBREAK_REDIRECT_INPUT_STRING,
		WORDBREAK_REDIRECT_INPUT_DUPLICATE,
		WORDBREAK_REDIRECT_INPUT_OUTPUT:
		return true
	default:
		return false
	}

}

func wordbreakType(t Token) WordbreakType {
	switch t.RawValue {
	case "<":
		return WORDBREAK_REDIRECT_INPUT
	case ">":
		return WORDBREAK_REDIRECT_OUTPUT
	case ">>":
		return WORDBREAK_REDIRECT_OUTPUT_APPEND
	case "&>", ">&":
		return WORDBREAK_REDIRECT_OUTPUT_BOTH
	case "&>>":
		return WORDBREAK_REDIRECT_OUTPUT_BOTH_APPEND
	case "<<<":
		return WORDBREAK_REDIRECT_INPUT_STRING
	case "<&":
		return WORDBREAK_REDIRECT_INPUT_DUPLICATE
	case "<>":
		return WORDBREAK_REDIRECT_INPUT_OUTPUT
	case "|":
		return WORDBREAK_PIPE
	case "|&":
		return WORDBREAK_PIPE_WITH_STDERR
	case "&":
		return WORDBREAK_LIST_ASYNC
	case ";":
		return WORDBREAK_LIST_SEQUENTIAL
	case "&&":
		return WORDBREAK_LIST_AND
	case "||":
		return WORDBREAK_LIST_OR
	default:
		// TODO check COMP_WORDBREAKS -> WORDBREAK_OTHER
		return WORDBREAK_UNKNOWN
	}
}
