package cyflag

import (
	"fmt"
	"github.com/yah01/cyds"
	"strconv"
	"strings"
)

type Parser struct {
	trie  cyds.Trie
	flags []cyflag

	LeftArgs []string
}

func (parser *Parser) bind(v interface{}, name string, defaultValue interface{}, usage string) {
	flag := &cyflag{
		variable:     v,
		defaultValue: defaultValue,
		Name:         name,
		Usage:        usage,
	}
	parser.trie.Insert(name, flag)
	parser.flags = append(parser.flags, *flag)
}

//func (parser *Parser) Bind(v interface{}, name string, defaultValue interface{}, usage string) {
//	switch v.(type) {
//	case *bool, *int, *string:
//		parser.bind(v, name, defaultValue, usage)
//	}
//}

func (parser *Parser) BoolVar(v *bool, name string, defaultValue bool, usage string) {
	parser.bind(v, name, defaultValue, usage)
}

func (parser *Parser) IntVar(v *int, name string, defaultValue int, usage string) {
	parser.bind(v, name, defaultValue, usage)
}

func (parser *Parser) FloatVar(v *float64, name string, defaultValue float64, usage string) {
	parser.bind(v, name, defaultValue, usage)
}

func (parser *Parser) StringVar(v *string, name string, defaultValue string, usage string) {
	parser.bind(v, name, defaultValue, usage)
}

func (parser *Parser) Parse(args []string) error {
	for i, _ := range parser.flags {
		flag := &parser.flags[i]
		switch flag.variable.(type) {
		case *bool:
			*(flag.variable.(*bool)) = flag.defaultValue.(bool)
		case *int:
			*(flag.variable.(*int)) = flag.defaultValue.(int)
		case *float64:
			*(flag.variable.(*float64)) = flag.defaultValue.(float64)
		case *string:
			*(flag.variable.(*string)) = flag.defaultValue.(string)
		}
	}

	for i := 0; i < len(args); i++ {
		s := args[i]
		if node, ok := parser.trie.Match(s); ok {
			for i := 0; i < len(node.Values); i++ {
				flag := node.Values[i].(*cyflag)

				switch flag.variable.(type) {
				case *bool:
					*(flag.variable.(*bool)) = true

				case *int:
					if i+1 >= len(args) {
						return fmt.Errorf("Miss interger, usage: %v", flag.Usage)
					}
					var err error
					*(flag.variable.(*int)), err = strconv.Atoi(args[i+1])
					if err != nil {
						return err
					}
					i++

				case *float64:
					if i+1 >= len(args) {
						return fmt.Errorf("Miss decimal, usage: %v", flag.Usage)
					}
					var err error
					*(flag.variable.(*float64)), err = strconv.ParseFloat(args[i+1], 64)
					if err != nil {
						return err
					}
					i++

				case *string:
					if i+1 >= len(args) {
						return fmt.Errorf("Miss string, usage: %v", flag.Usage)
					}
					*(flag.variable.(*string)) = args[i+1]
					i++
				}
			}
		} else {
			parser.LeftArgs = append(parser.LeftArgs, s)
		}
	}

	return nil
}

func (parser *Parser) ParseString(str string) error {
	return parser.Parse(strings.Fields(str))
}

func (parser *Parser) Clear() {
	*parser = Parser{}
}
