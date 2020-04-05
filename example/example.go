package main

import (
	"fmt"
	"github.com/yah01/cyflag"
	"log"
)

func main() {
	var (
		parser cyflag.Parser
		args   = "i love cyflag -best -times 95"

		best  bool
		times int
		love  string
	)
	//parser.BoolVar(&best, "-best", false, "whether the best")
	//parser.IntVar(&times, "-times", 0, "-times [int]")
	//parser.StringVar(&love, "love", "something", "love [string]")

	parser.Bind(&best, "-best", false, "whether the best")
	parser.Bind(&times, "-times", 0, "-times [int]")
	parser.Bind(&love, "love", "something", "love [string]")

	if err := parser.ParseString(args); err != nil {
		log.Println(err)
	}

	fmt.Printf("best: %+v\n"+
		"times: %+v\n"+
		"love: %+v\n",
		best, times, love)
}
