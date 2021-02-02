package main

import (
  "testing"
  "net/http"
  "net/http/httptest"
)

func TestRoot(t *testing.T) {
  req := httptest.NewRequest("GET", "/somepath", nil)
  rr := httptest.NewRecorder()

  handler := http.HandlerFunc(rootHandler)
  handler.ServeHTTP(rr, req)

  if status := rr.Code; status != http.StatusOK {
    t.Errorf("Status code is incorrect, got '%v', expected '%v'",
      status, http.StatusOK)
  }

  expected := "Path: somepath"

  if rr.Body.String() != expected {
    t.Errorf("Handler error, got '%v', expected '%v'",
      rr.Body.String(), expected)
  }
} 
