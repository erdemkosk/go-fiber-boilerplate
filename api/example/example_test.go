package example

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestGetExamples(t *testing.T) {
	app := fiber.New()
	app.Get("/api/examples", GetExamples)

	req := httptest.NewRequest(http.MethodGet, "/api/examples", nil)
	resp, err := app.Test(req, -1)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var examples []Example
	err = json.NewDecoder(resp.Body).Decode(&examples)
	assert.NoError(t, err)
	assert.NotNil(t, examples)
}

func TestCreateExample(t *testing.T) {
	app := fiber.New()
	app.Post("/api/examples", CreateExample)

	payload := `{"name":"John Doe"}`
	req := httptest.NewRequest(http.MethodPost, "/api/examples", strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req, -1)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var example Example
	err = json.NewDecoder(resp.Body).Decode(&example)
	assert.NoError(t, err)
	assert.NotNil(t, example)
	assert.Equal(t, "John Doe", example.Name)
}
