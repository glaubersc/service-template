package integration

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/your-org/service-template/internal/interfaces/rest"
)

func TestHealthEndpoint(t *testing.T) {
	router := rest.NewRouter()
	server := httptest.NewServer(router)
	defer server.Close()

	resp, err := http.Get(server.URL + "/health")
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200, got %d", resp.StatusCode)
	}
}
