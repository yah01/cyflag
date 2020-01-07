package cyFlag

import (
	"cyDS"
	"errors"
	"os"
	"strconv"
	"strings"
)

var (
	trie cyDS.Trie
	Args []string
)

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
	Args = make([]string, 1)
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
					return errors.New("no string value")
				}

				var err error
				*(node.Value.(*int)), err = strconv.Atoi(args[i+1])
				if err != nil {
					return err
				}
				i++
			default:
				Args = append(Args, s)
			}

		}
	}

	return nil
}

func ParseString(str string) error {
	Args = make([]string, 1)
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
					return errors.New("no string value")
				}

				var err error
				*(node.Value.(*int)), err = strconv.Atoi(args[i+1])
				if err != nil {
					return err
				}
				i++
			default:
				Args = append(Args, s)
			}

		}
	}

	return nil
}
