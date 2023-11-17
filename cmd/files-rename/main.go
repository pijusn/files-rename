package main

import (
	"flag"
	"log"

	"github.com/pijusn/files-rename/lib"
)

func parsedFlags() *lib.Config {
	config := &lib.Config{}

	flag.StringVar(&config.Directory, "directory", ".", "path to a directory to scan for files")
	flag.StringVar(&config.Name, "name", "%4d", "pattern defining new file name")
	flag.Parse()

	return config
}

func main() {
	config := parsedFlags()
	err := lib.Run(config)
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}
}
