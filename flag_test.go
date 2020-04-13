package cyflag

import (
	"testing"
)

func TestFlag(t *testing.T) {
	var (
		parser Parser
		str    = `i love "cyber flag" -best`

		best          bool
		loveSomething string
	)
	t.Log(Args)
	parser.BoolVar(&best, "-best", false, "-best")
	parser.StringVar(&loveSomething, "love", "something", "love [string]")

	parser.ParseString(str)

	t.Log(best, loveSomething)
}

func TestSplitArgsString(t *testing.T) {
	args := ` test split args = "hello world" args2  = ' "cyber" "flag" '`
	argList := splitArgsString(args)

	for i,arg := range argList {
		t.Log(i,arg)
	}
}