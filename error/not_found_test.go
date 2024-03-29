package error_test

import (
	"eve.vehicle.api.com/m/v2/error"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNotFoundError(t *testing.T) {
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	error.NotFoundError(c)

	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Contains(t, w.Body.String(), `"error":"NOTFOUNDERR-1"`)
	assert.Contains(t, w.Body.String(), `"message":"Vehicle not found"`)
}
