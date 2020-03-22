package cyflag

import (
	"testing"
)

func TestFlag(t *testing.T) {
	var (
		parser Parser
		str    = "i love cyflag -best"

		best          bool
		loveSomething string
	)
	parser.BoolVar(&best, "-best", false, "-best")
	parser.StringVar(&loveSomething, "love", "something", "love [string]")

	parser.ParseString(str)

	t.Log(best, loveSomething)
}
