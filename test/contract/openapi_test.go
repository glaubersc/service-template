package contract

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestOpenAPISpecIsValid(t *testing.T) {
	loader := openapi3.NewLoader()
	doc, err := loader.LoadFromFile("../../api/openapi/openapi.yaml")
	if err != nil {
		t.Fatalf("failed to load openapi spec: %v", err)
	}

	if err := doc.Validate(loader.Context); err != nil {
		t.Fatalf("invalid openapi spec: %v", err)
	}
}
