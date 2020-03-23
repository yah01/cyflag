package cyflag

import (
	"os"
)

type cyflag struct {
	variable     interface{}
	defaultValue interface{}

	Name  string
	Usage string
}

var (
	osParser Parser
)

func BoolVar(v *bool, name string, defaultValue bool, usage string) {
	osParser.BoolVar(v, name, defaultValue, usage)
}

func IntVar(v *int, name string, defaultValue int, usage string) {
	osParser.IntVar(v, name, defaultValue, usage)
}

func FloatVar(v *float64, name string, defaultValue float64, usage string) {
	osParser.bind(v, name, defaultValue, usage)
}

func StringVar(v *string, name string, defaultValue string, usage string) {
	osParser.StringVar(v, name, defaultValue, usage)
}

func Parse() error {
	return osParser.Parse(os.Args[1:])
}

func Clear() {
	osParser.Clear()
}

func Usage() {

}
