package main

import (
	"github.com/linuxfight/yandexCalcApi/cmd/application"
	"net/http"
	"strings"
	"testing"
)

func TestServer(t *testing.T) {
	tests := []struct {
		method    string
		route     string
		status    int
		body      string
		setHeader bool
	}{
		{
			method: "GET",
			route:  "/",
			status: 404,
		},
		{
			method: "GET",
			route:  "/api/v1/calculate",
			status: 405,
		},
		{
			method:    "POST",
			route:     "/api/v1/calculate",
			status:    422,
			body:      "{",
			setHeader: true,
		},
		{
			method:    "POST",
			route:     "/api/v1/calculate",
			status:    422,
			body:      "{ \"expression\": \"2+\" }",
			setHeader: true,
		},
		{
			method:    "POST",
			route:     "/api/v1/calculate",
			status:    422,
			body:      "{ \"expression\": \"2+2\" }",
			setHeader: false,
		},
		{
			method:    "POST",
			route:     "/api/v1/calculate",
			status:    200,
			body:      "{ \"expression\": \"2+2\" }",
			setHeader: true,
		},
	}

	app := application.New()

	for _, test := range tests {
		req, err := http.NewRequest(test.method, test.route, strings.NewReader(test.body))
		if err != nil {
			t.Fatal(err)
		}
		if test.setHeader {
			req.Header.Set("Content-Type", "application/json")
		}
		res, err := app.Http.Test(req, -1)
		if err != nil {
			t.Fatal(err)
		}
		if res.StatusCode != test.status {
			t.Fatalf("Test failed - %d; want %d", res.StatusCode, test.status)
		}
	}
}
