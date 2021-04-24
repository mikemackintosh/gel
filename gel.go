package gel

import (
	"flag"
	"fmt"
	"os"
	"reflect"
	"time"
)

var (
	// Registry stores the active configuration.
	registry = NewRegistry()

	// set debug
	Debug bool = false
)

// This keeps track of our different source types.
const (
	Env int = iota
	Config
	Flags
)

// ConfigFile is a filename as a string.
type ConfigFile string

// Registry defines the configuraiton.
type Registry struct {
	Config    *ConfigFile
	LoadOrder []int
	Items     map[string]*Item
}

// NewRegistry creates a fresh registry.
func NewRegistry() *Registry {
	// Lets reset the stdlib flag set.
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ContinueOnError)

	// Create a new registry.
	return &Registry{
		Items:     map[string]*Item{},
		LoadOrder: []int{Env, Config, Flags},
	}
}

// Item defines the structure of a configuration entry.
type Item struct {
	t           reflect.Kind
	Name        string
	Description string
	Default     interface{}
	Value       interface{}
}

// Int casts the interface{} value as an integer.
func (i *Item) Int() int {
	return i.Value.(int)
}

// Bool casts the interface{} value as an Bool.
func (i *Item) Bool() bool {
	return i.Value.(bool)
}

// Duration casts the interface{} value as an Duration.
func (i *Item) Duration() time.Duration {
	return i.Value.(time.Duration)
}

// String casts the interface{} value as an String.
func (i *Item) String() string {
	if i.Value != nil {
		return i.Value.(string)
	}
	return ""
}


func Bool(name string, defaultValue bool, description string) {
	(*registry).Items[name] = &Item{
		t:           reflect.Bool,
		Name:        name,
		Description: description,
		Value:     defaultValue,
	}
}

func Int(name string, defaultValue int, description string) {
	(*registry).Items[name] = &Item{
		t:           reflect.Int,
		Name:        name,
		Description: description,
		Value:     defaultValue,
	}
}

func String(name, defaultValue, description string) {
	(*registry).Items[name] = &Item{
		t:           reflect.String,
		Name:        name,
		Description: description,
		Value:       defaultValue,
		Default: 		 defaultValue,
	}
}

func Duration(name string, defaultValue string, description string) {
	d, _ := time.ParseDuration(defaultValue)
	(*registry).Items[name] = &Item{
		t:           reflect.Int64,
		Name:        name,
		Description: description,
		Value:     d,
	}
}

func Get(name string) (*Item, error) {
	if v, ok := (*registry).Items[name]; ok {
		return v, nil
	}
	return nil, fmt.Errorf("no matching entry for '%s'", name)
}

// MustGet ignores any errors and returns
func MustGet(name string) *Item {
	if v, ok := (*registry).Items[name]; ok {
		return v
	}
	return &Item{}
}

// parseFlags will read the flags provided to the binary.
// this section is very inefficient and needs to be readdressed.
func parseFlags() {

	for n, i := range (*registry).Items {
		switch i.t {
		case reflect.String:
			var v string
			flag.StringVar(&v, n, i.Default.(string), i.Description)
			flag.Parse()
			i.Value = v

		case reflect.Int:
			var v int
			flag.IntVar(&v, n, i.Value.(int), i.Description)
			flag.Parse()
			i.Value = v

		case reflect.Bool:
			var v bool
			flag.BoolVar(&v, n, i.Value.(bool), i.Description)
			flag.Parse()
			i.Value = v

		case reflect.Int64:
			var v time.Duration
			flag.DurationVar(&v, n, i.Value.(time.Duration), i.Description)
			flag.Parse()
			i.Value = v
		}
	}
}

// parseEnv reads the environment variables and creates a registry item.
func parseEnv() {
	for n, i := range (*registry).Items {
		env := os.Getenv(n)
		if len(env) > 0 {
			i.Value = env
		}
	}
}

func parseConfig() {
	if (*registry).Config != nil {
		// TODO
	}
}

// Up loads the flags and configuration and sets the values accordingly.
func Up() {
	for i := len((*registry).LoadOrder) - 1; i >= 0; i-- {
		switch (*registry).LoadOrder[i] {
		case Env:
			parseEnv()
		case Config:
			parseConfig()
		case Flags:
			parseFlags()
		}
	}
}

// UseOrder sets the order of preference.
func UseOrder(order ...int) {
	registry.LoadOrder = order
}

// UseConfig overrides any configuration file passed.
func UseConfig(file string, to *interface{}) error {
	c := ConfigFile(file)
	registry.Config = &c
	return nil
}
