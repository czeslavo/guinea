package guinea

import (
	"flag"
)

// OptionValue stores the value of a parsed option as returned by the standard
// library flag package. The helper methods can be used to cast the value
// quickly but they will only succeed if the defined type of the option matches
// the called method.
type OptionValue struct {
	Value interface{}
}

// Bool casts a value to a bool and panicks on failure.
func (v OptionValue) Bool() bool {
	return *v.Value.(*bool)
}

// Int casts a value to an int and panicks on failure.
func (v OptionValue) Int() int {
	return *v.Value.(*int)
}

// Str casts a value to a string and panicks on failure.
func (v OptionValue) Str() string {
	return *v.Value.(*string)
}

// Context holds the options and arguments provided by the user.
type Context struct {
	Options   map[string]OptionValue
	Arguments []string
}

func makeContext(c Command, args []string) (*Context, error) {
	context := &Context{
		Options: make(map[string]OptionValue),
	}

	flagset := flag.NewFlagSet("sth", flag.ContinueOnError)
	flagset.Usage = func() {}
	for _, option := range c.Options {
		switch option.Type {
		case String:
			if option.Default == nil {
				option.Default = ""
			}
			context.Options[option.Name] = OptionValue{
				Value: flagset.String(option.Name, option.Default.(string), ""),
			}
		case Bool:
			if option.Default == nil {
				option.Default = false
			}
			context.Options[option.Name] = OptionValue{
				Value: flagset.Bool(option.Name, option.Default.(bool), ""),
			}
		case Int:
			if option.Default == nil {
				option.Default = 0
			}
			context.Options[option.Name] = OptionValue{
				Value: flagset.Int(option.Name, option.Default.(int), ""),
			}
		}
	}
	e := flagset.Parse(args)
	if e != nil {
		return nil, e
	}
	context.Arguments = flagset.Args()
	return context, nil
}
