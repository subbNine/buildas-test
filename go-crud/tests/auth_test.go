package tests

import (
    "net/http"
    "net/http/httptest"
    "testing"
    "go-crud/internal/handlers"
    "strings"
)

func TestRegister(t *testing.T) {
    payload := `{"username": "testuser", "password": "password"}`
    req, _ := http.NewRequest("POST", "/register", strings.NewReader(payload))
    w := httptest.NewRecorder()

    handlers.Register(w, req)

    if status := w.Code; status != http.StatusOK {
        t.Errorf("expected status %d, got %d", http.StatusOK, status)
    }
}
