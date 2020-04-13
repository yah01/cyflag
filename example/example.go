package main

import (
	"fmt"
	"github.com/yah01/cyflag"
)

func main() {
	var (
		parser cyflag.Parser
		args   = `i love "cyber flag" -best -times 95`

		best  bool
		times int
		love  string
	)

	parser.Bind(&best, "-best", false, "whether the best")
	parser.Bind(&times, "-times", 0, "-times [int]")
	parser.Bind(&love, "love", "something", "love [string]")

	parser.ParseString(args)

	fmt.Printf("best: %+v\n"+
		"times: %+v\n"+
		"love: %+v\n",
		best, times, love)
}
