package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	req := httptest.NewRequest("GET", "/cafe?count=10&city=moscow", nil) // здесь нужно создать запрос к сервису

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)
	status := responseRecorder.Code

	require.Equal(t, http.StatusOK, status)

	body := responseRecorder.Body.String()
	list := strings.Split(body, ",")

	assert.Len(t, list, totalCount)
	// здесь нужно добавить необходимые проверки
}

func TestMainHandlerAllOK(t *testing.T) {
	//totalCount := 4
	req := httptest.NewRequest("GET", "/cafe?count=10&city=moscow", nil) // здесь нужно создать запрос к сервису

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)
	status := responseRecorder.Code

	require.Equal(t, http.StatusOK, status)

	assert.NotEmpty(t, responseRecorder.Body)
	// здесь нужно добавить необходимые проверки
}

func TestMainHandlerCityNotAllowed(t *testing.T) {
	bodyres := "wrong city value"
	req := httptest.NewRequest("GET", "/cafe?count=10&city=spb", nil) // здесь нужно создать запрос к сервису

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	status := responseRecorder.Code

	require.Equal(t, http.StatusBadRequest, status)

	body := responseRecorder.Body.String()

	assert.Equal(t, bodyres, body)
	// здесь нужно добавить необходимые проверки
}
