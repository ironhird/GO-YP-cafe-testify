package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Запрос сформирован корректно, сервис возвращает код ответа 200 и тело ответа не пустое
func TestMainHandlerWhenOk(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=2&city=moscow", nil)
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, responseRecorder.Code, http.StatusOK) // код ответа 200
	assert.NotEmpty(t, responseRecorder.Body)              // тело ответа не пустое
}

// Сервис возвращает код ответа 400 и ошибку "wrong city value" в теле ответа
func TestMainHandlerWhenWrongCity(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=1&city=penza", nil)
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, responseRecorder.Code, http.StatusBadRequest)      // код ответа 400
	assert.Equal(t, responseRecorder.Body.String(), "wrong city value") // "wrong city value" в теле ответа
}

// Если в параметре count указано больше, чем есть всего, должны вернуться все доступные кафе.
func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=50&city=moscow", nil)
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	bodyString := responseRecorder.Body.String()
	cafeCount := len(strings.Split(bodyString, ","))
	actualCafeCount := len(cafeList["moscow"])

	require.Equal(t, responseRecorder.Code, http.StatusOK) // код ответа 200
	assert.Equal(t, actualCafeCount, cafeCount)            // все доступные кафе
}
