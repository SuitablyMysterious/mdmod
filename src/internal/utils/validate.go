package utils

import (
	"log"
	"os"

	"github.com/kaptinlin/jsonschema"
)

func validate(mdspecPath string, schemaPath string) bool {

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
	mdspec, err := os.Open(mdspecPath)
	if err != nil {
		log.Fatalf("Failed to open %s", mdspecPath)
	} else {
		log.Printf("Opened %s", mdspecPath)
	}

	// Get the schema file
	compiler := jsonschema.NewCompiler() // Create a JSON compiler for the schema
	openedSchema, err := os.ReadFile(schemaPath)
	if err != nil {
		log.Fatalf("Failed to read schema: %s", schemaPath)
	} else {
		log.Printf("Read schema: %s", schemaPath)
	}
	schema, err := compiler.Compile(openedSchema)
	if err != nil {
		log.Fatal("Failed to compile schema")
	}

	// Return the answer
	return schema.Validate(mdspec).IsValid()
}
