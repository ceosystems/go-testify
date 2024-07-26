
package main

import (
    "net/http"
    "net/http/httptest"
   // "strconv"
    "strings"
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestCorrectRequestReturns200AndNonEmptyBody(t *testing.T) {
    req := httptest.NewRequest("GET", "/cafe?count=2&city=moscow", nil)
    rec := httptest.NewRecorder()

    mainHandle(rec, req)

    assert.Equal(t, rec.Code, http.StatusOK)
    assert.NotEmpty(t, rec.Body.String())
}

func TestUnsupportedCityValueReturns400AndErrorMessage(t *testing.T) {
    req := httptest.NewRequest("GET", "/cafe?count=2&city=unsupportedCity", nil)
    rec := httptest.NewRecorder()

    mainHandle(rec, req)

    assert.Equal(t, rec.Code, http.StatusBadRequest)
    assert.Contains(t, rec.Body.String(), "wrong city value")
}

func TestCountGreaterThanTotalReturnsAllCafes(t *testing.T) {
    req := httptest.NewRequest("GET", "/cafe?count=5&city=moscow", nil)
    rec := httptest.NewRecorder()

    mainHandle(rec, req)

    assert.Equal(t, rec.Code, http.StatusOK)
    assert.Len(t, strings.Split(rec.Body.String(), ","), 4)
}