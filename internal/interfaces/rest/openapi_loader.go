package rest

import (
	"log"

	"github.com/getkin/kin-openapi/openapi3"
)

func LoadOpenAPI(path string) *openapi3.T {
	loader := openapi3.NewLoader()
	doc, err := loader.LoadFromFile(path)
	if err != nil {
		log.Fatalf("failed to load openapi spec: %v", err)
	}

	if err := doc.Validate(loader.Context); err != nil {
		log.Fatalf("invalid openapi spec: %v", err)
	}

	return doc
}
