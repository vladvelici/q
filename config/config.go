package config

import (
	"github.com/namsral/flag"
)

// Config is the struct holding all configuration for running the server and
// client.
type Config struct {
	Address string
}

// Defaults is the default config for the q project.
var Defaults = Config{
	Address: "127.0.0.1:3399",
}

// Reader holds a FlagSet setup to read required parameters from the command
// line or environment variables (prefixed with "Q_"). After the FlagSet has
// been parsed you can call reader.MakeConfig() to generate a Config struct.
type Reader struct {
	Flagset *flag.FlagSet

	address *string
}

// NewReader creates a new reader with the given name.
func NewReader(name string) Reader {
	f := flag.NewFlagSetWithEnvPrefix(name, "Q_", 0)

	r := Reader{
		Flagset: f,
		address: f.String("address", "127.0.0.1:3399", "Address the daemon listens to, defaults to 127.0.0.1:3399"),
	}

	return r
}

// MakeConfig generates a config using the parsed flagset. Do not call before parsing the flagset.
func (r Reader) MakeConfig() Config {
	return Config{
		Address: *r.address,
	}
}
