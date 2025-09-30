package spec

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

func main() {
	log.SetOutput(os.Stdout) // will write to a file later
	log.SetPrefix("logger.go: ")
	yaml.NewDecoder(nil) // just to use the package
}