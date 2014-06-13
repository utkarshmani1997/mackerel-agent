package config

import "path/filepath"
import "os"

var mackerelRoot = filepath.Join(os.Getenv("HOME"), "Library", "mackerel-agent")

// The default configuration for dawrin.
var DefaultConfig = &Config{
	Apibase:  "https://mackerel.io",
	Root:     mackerelRoot,
	Pidfile:  filepath.Join(mackerelRoot, "pid"),
	Conffile: filepath.Join(mackerelRoot, "mackerel-agent.conf"),
	Roles:    []string{},
	Verbose:  false,
}
