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
	Args     []string
)

func init() {
	copy(Args, os.Args)
}

func BoolVar(v *bool, name string, defaultValue bool, usage string) {
	osParser.BoolVar(v, name, defaultValue, usage)
}

func IntVar(v *int, name string, defaultValue int, usage string) {
	osParser.IntVar(v, name, defaultValue, usage)
}

func UintVar(v *uint, name string, defaultValue uint, usage string) {
	osParser.UintVar(v, name, defaultValue, usage)
}

func FloatVar(v *float64, name string, defaultValue float64, usage string) {
	osParser.bind(v, name, defaultValue, usage)
}

func StringVar(v *string, name string, defaultValue string, usage string) {
	osParser.StringVar(v, name, defaultValue, usage)
}

func Parse() error {
	err := osParser.Parse(os.Args[1:])
	Args = osParser.LeftArgs
	return err
}

func PrintUsage() {
	osParser.PrintUsage()
}

func Clear() {
	osParser.Clear()
}
