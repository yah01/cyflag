package cyFlag

import (
	"fmt"
	"testing"
)

var (
	b bool
	i int
	s string
)

func testCase(args string) {
	b = false
	i = 57
	s = "default"
	err := ParseString(args)
	fmt.Println(b, i, s)
	if err != nil {
		fmt.Println(err)
	}
}
func TestFlag(t *testing.T) {

	BoolVar(&b, "-bool", false, "test bool")
	IntVar(&i, "-int", 57, "test int")
	StringVar(&s, "-str", "default", "test string")

	testCase("-bool -int 10 -str hello")
	testCase("-bool -int")
	testCase("-int 123a -str 231")
	testCase("-int 235 -str zxc -bool")
}
