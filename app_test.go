package main

import (
	"io"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2/utils"
)

func TestServer(t *testing.T) {
	app := NewServer()
	resp, err := app.Test(httptest.NewRequest("GET", "/", nil))

	defer resp.Body.Close()
	bodyBytes, err := io.ReadAll(resp.Body)

	wantBody := "Server is running! Send your request"
	gotBody := string(bodyBytes)

	utils.AssertEqual(t, nil, err, "App test")
	utils.AssertEqual(t, 200, resp.StatusCode, "Status code")

	if gotBody != wantBody {
		t.Errorf("body got %v got %v", gotBody, wantBody)
	}
}
