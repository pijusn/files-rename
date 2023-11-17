package lib

import "fmt"

type Config struct {
	Directory string
	Name      string
}

func (config *Config) Validate() error {
	if config == nil {
		return fmt.Errorf("no configuration given")
	}
	if config.Directory == "" {
		return fmt.Errorf("directory is not specified")
	}
	if config.Name == "" {
		return fmt.Errorf("name pattern is not specified")
	}
	return nil
}
