package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetBookRouting(t *testing.T) {
	router := gin.Default()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/book/api/v1/", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, "201")
}
