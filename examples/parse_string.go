package examples

import (
	"fmt"
	"github.com/yah01/cyflag"
)

func ParseString() {
	var (
		parser cyflag.Parser
		str    = "i love cyflag -best"

		best          bool
		loveSomething string
	)
	parser.BoolVar(&best, "-best", false, "-best")
	parser.StringVar(&loveSomething, "love", "something", "love [string]")

	parser.ParseString(str)

	fmt.Println(best,loveSomething)
}