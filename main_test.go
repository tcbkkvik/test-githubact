package main

import (
    "net/http"
    "net/http/httptest"
    "testing"
)

func TestHelloWorld(t *testing.T) {
    req, err := http.NewRequest("GET", "/", nil)
    if err != nil {
        t.Fatal(err)
    }

    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(helloWorld)

    handler.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusOK {
        t.Errorf("Expected status code 200, got %v", status)
    }

    expected := "Hello, World!"
    if rr.Body.String() != expected {
        t.Errorf("Expected body %v, got %v", expected, rr.Body.String())
    }
}