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
