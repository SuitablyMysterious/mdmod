package spec
import (
	"log"
	"os"
	"github.com/kaptinlin/jsonschema"
)

func validate(mdspec_path string, schema_path string) (bool) {

	// Introduce logging
	logfile, err := os.Open("../log.log")
	if err != nil {
		log.Fatal("Failed to open log file")
	} else {
		log.Print("Opened log file")
	}
	log.SetOutput(logfile)
	log.SetPrefix("logger.go: ")

	// Open the .mdspec file
	mdspec, err := os.Open(mdspec_path)
	if err != nil {
		log.Fatalf("Failed to open %s", mdspec_path)
	} else {
		log.Printf("Opened %s", mdspec_path)
	}
	
	// Get the schema file
	compiler := jsonschema.NewCompiler() // Create a JSON compiler for the schema
	opened_schema, err := os.ReadFile(schema_path)
	if err != nil {
		log.Fatalf("Failed to read schema: %s", schema_path)
	} else {
		log.Printf("Read schema: %s", schema_path)
	}
	schema, err := compiler.Compile(opened_schema)
	if err != nil {
		log.Fatal("Failed to compile schema")
	}

	// Return the answer
	return schema.Validate(mdspec).IsValid()
}