package cyflag

import (
	"fmt"
	"github.com/yah01/cyds"
	"reflect"
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

func (parser *Parser) Bind(v interface{}, name string, defaultValue interface{}, usage string) {

	switch v.(type) {
	case *bool:
		parser.bind(v, name, reflect.ValueOf(defaultValue).Bool(), usage)
	case *int, *int8, *int16, *int32, *int64:
		parser.bind(v, name, reflect.ValueOf(defaultValue).Int(), usage)
	case *uint, *uint8, *uint16, *uint32, *uint64:
		parser.bind(v, name, reflect.ValueOf(defaultValue).Uint(), usage)
	case *float32, *float64:
		parser.bind(v, name, reflect.ValueOf(defaultValue).Float(), usage)
	case *string:
		if kind := reflect.TypeOf(defaultValue).Kind(); kind == reflect.String {
			parser.bind(v, name, defaultValue, usage)
		} else {
			panic(reflect.ValueError{
				Method: "reflect.Value.String",
				Kind:   kind,
			})
		}
	}
}

func (parser *Parser) BoolVar(v *bool, name string, defaultValue bool, usage string) {
	parser.bind(v, name, defaultValue, usage)
}

func (parser *Parser) IntVar(v *int, name string, defaultValue int, usage string) {
	parser.bind(v, name, defaultValue, usage)
}

func (parser *Parser) UintVar(v *uint, name string, defaultValue uint, usage string) {
	parser.bind(v, name, defaultValue, usage)
}

func (parser *Parser) FloatVar(v *float64, name string, defaultValue float64, usage string) {
	parser.bind(v, name, defaultValue, usage)
}

func (parser *Parser) StringVar(v *string, name string, defaultValue string, usage string) {
	parser.bind(v, name, defaultValue, usage)
}

func (parser *Parser) Parse(args []string) error {
	if args == nil {
		args = parser.LeftArgs
	} else {
		parser.LeftArgs = nil
	}
	for _, flag := range parser.flags {
		refVar := reflect.ValueOf(flag.variable).Elem()
		defaultValue := reflect.ValueOf(flag.defaultValue)
		switch flag.variable.(type) {
		case *bool:
			*(flag.variable.(*bool)) = flag.defaultValue.(bool)
		case *int, *int8, *int16, *int32, *int64:
			refVar.SetInt(defaultValue.Int())
		case *uint, *uint8, *uint16, *uint32, *uint64:
			refVar.SetUint(defaultValue.Uint())
		case *float32, *float64:
			refVar.SetFloat(defaultValue.Float())
		case *string:
			*(flag.variable.(*string)) = flag.defaultValue.(string)
		}
	}

	for i := 0; i < len(args); i++ {
		s := args[i]
		if node, ok := parser.trie.Match(s); ok {
			for j := 0; j < len(node.Values); j++ {
				flag := node.Values[j].(*cyflag)
				refVar := reflect.ValueOf(flag.variable).Elem()

				switch flag.variable.(type) {
				case *bool:
					*(flag.variable.(*bool)) = true

				case *int, *int8, *int16, *int32, *int64:
					if i+1 >= len(args) {
						return fmt.Errorf("Miss interger, usage: %v", flag.Usage)
					}

					value, err := strconv.ParseInt(args[i+1], 10, 64)
					if err != nil {
						return err
					}
					refVar.SetInt(value)
					i++

				case *uint, *uint8, *uint16, *uint32, *uint64:
					if i+1 >= len(args) {
						return fmt.Errorf("Miss interger, usage: %v", flag.Usage)
					}

					value, err := strconv.ParseUint(args[i+1], 10, 64)
					if err != nil {
						return err
					}

					refVar.SetUint(value)
					i++

				case *float32, *float64:
					if i+1 >= len(args) {
						return fmt.Errorf("Miss decimal, usage: %v", flag.Usage)
					}
					value, err := strconv.ParseFloat(args[i+1], 64)
					if err != nil {
						return err
					}
					refVar.SetFloat(value)
					i++

				case *string:
					if i+1 >= len(args) {
						return fmt.Errorf("Miss string, usage: %v", flag.Usage)
					}

					var offset int
					*(flag.variable.(*string)) = args[i+1]
					i += offset
				}
			}
		} else {
			parser.LeftArgs = append(parser.LeftArgs, s)
		}
	}

	return nil
}

func (parser *Parser) ParseString(str string) error {
	return parser.Parse(splitArgsString(str))
}

func (parser *Parser) PrintUsage() {
	for _, flag := range parser.flags {
		fmt.Printf("%s: %s\n", flag.Name, flag.Usage)
	}
}

func (parser *Parser) Clear() {
	*parser = Parser{}
}

// Split arguments string into string slice.
func splitArgsString(args string) []string {
	args = strings.TrimSpace(args)
	args += " "
	var (
		args_rune          = []rune(args)
		leftQuote          rune
		leftQuotePosition  int
		hasLeftQuote       bool
		left2RightQuoteMap = make(map[int]int)
		argList            []string
	)

	// Find all quote pairs
	for i := 0; i < len(args_rune); i++ {
		c := args_rune[i]

		// For string argument with space
		if c == '\'' || c == '"' {
			if hasLeftQuote && c == leftQuote {
				left2RightQuoteMap[leftQuotePosition] = i
				hasLeftQuote = false
			} else if !hasLeftQuote {
				leftQuote = c
				hasLeftQuote = true
				leftQuotePosition = i
			}
		}
	}

	var lastPosition int
	for i := 0; i < len(args_rune); i++ {
		c := args_rune[i]

		var (
			arg              string
			allowEmptyString bool
		)

		switch c {
		case '\'', '"':
			if rightQuotePosition, ok := left2RightQuoteMap[i]; ok {
				allowEmptyString = true
				arg = strings.Trim(string(args_rune[i:rightQuotePosition]), string(c))
				lastPosition = rightQuotePosition + 1
				i = rightQuotePosition
			}
		case ' ':
			arg = string(args_rune[lastPosition:i])
			lastPosition = i + 1
		default:
			continue
		}

		if len(arg) > 0 || allowEmptyString {
			argList = append(argList, arg)
		}
	}

	return argList
}