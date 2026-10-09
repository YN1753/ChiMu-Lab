package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"chimu-lab/internal/config"

	"github.com/gin-gonic/gin"
)

func TestAdminAuthRequired(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("Passes when ADMIN_API_KEY is empty", func(t *testing.T) {
		config.AppConfig.AdminAPIKey = ""
		r := gin.New()
		r.Use(AdminAuthRequired())
		r.POST("/test", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
		})

		req := httptest.NewRequest(http.MethodPost, "/test", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		if w.Code != http.StatusOK {
			t.Errorf("expected 200, got %d", w.Code)
		}
	})

	t.Run("Rejects unauthorized request when ADMIN_API_KEY is configured", func(t *testing.T) {
		config.AppConfig.AdminAPIKey = "secret-test-token"
		r := gin.New()
		r.Use(AdminAuthRequired())
		r.POST("/test", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
		})

		// 1. 无 header
		req := httptest.NewRequest(http.MethodPost, "/test", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		if w.Code != http.StatusUnauthorized {
			t.Errorf("expected 401, got %d", w.Code)
		}

		// 2. 错误 token
		reqErr := httptest.NewRequest(http.MethodPost, "/test", nil)
		reqErr.Header.Set("Authorization", "Bearer wrong-token")
		wErr := httptest.NewRecorder()
		r.ServeHTTP(wErr, reqErr)
		if wErr.Code != http.StatusUnauthorized {
			t.Errorf("expected 401 for wrong token, got %d", wErr.Code)
		}

		// 3. 正确 token
		reqOk := httptest.NewRequest(http.MethodPost, "/test", nil)
		reqOk.Header.Set("Authorization", "Bearer secret-test-token")
		wOk := httptest.NewRecorder()
		r.ServeHTTP(wOk, reqOk)
		if wOk.Code != http.StatusOK {
			t.Errorf("expected 200 for correct token, got %d", wOk.Code)
		}
	})
}
