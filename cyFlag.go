package cyFlag

import (
	"cyDS"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type cyflag struct {
	Name  string
	Type  interface{}
	Usage string
}

var (
	trie  cyDS.Trie
	Args  []string
	flags []cyflag
)

func initData() {
	Args = make([]string, 0)
	flags = make([]cyflag, 0)
}

func BoolVar(v *bool, name string, defaultValue bool, usage string) {
	*v = defaultValue
	trie.Insert(name, v)
}

func IntVar(v *int, name string, defaultValue int, usage string) {
	*v = defaultValue
	trie.Insert(name, v)
}

func StringVar(v *string, name string, defaultValue string, usage string) {
	*v = defaultValue
	trie.Insert(name, v)
}

func Parse() error {
	initData()
	args := os.Args[1:]

	for i := 0; i < len(args); i++ {
		s := args[i]
		if node, ok := trie.Match(s); ok {
			switch node.Value.(type) {
			case *bool:
				*(node.Value.(*bool)) = true

			case *string:
				if i+1 >= len(args) {
					return errors.New("no string value")
				}

				*(node.Value.(*string)) = args[i+1]
				i++

			case *int:
				if i+1 >= len(args) {
					return errors.New("no int value")
				}

				var err error
				*(node.Value.(*int)), err = strconv.Atoi(args[i+1])
				if err != nil {
					return err
				}
				i++
			}

		} else {
			Args = append(Args, s)
		}
	}

	return nil
}

func ParseString(str string) error {
	initData()
	args := strings.Split(str, " ")

	for i := 0; i < len(args); i++ {
		s := args[i]
		if node, ok := trie.Match(s); ok {
			switch node.Value.(type) {
			case *bool:
				*(node.Value.(*bool)) = true

			case *string:
				if i+1 >= len(args) {
					return errors.New("no string value")
				}

				*(node.Value.(*string)) = args[i+1]
				i++

			case *int:
				if i+1 >= len(args) {
					return errors.New("no int value")
				}

				var err error
				*(node.Value.(*int)), err = strconv.Atoi(args[i+1])
				if err != nil {
					return err
				}
				i++
			}

		} else {
			Args = append(Args, s)
		}
	}

	return nil
}

func Usage() {
	for _, f := range flags {
		fmt.Print(f.Name, f.Type, "("+f.Usage+") ")
	}
}
