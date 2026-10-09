package middleware

import (
	"net/http"
	"strings"

	"chimu-lab/internal/config"

	"github.com/gin-gonic/gin"
)

// AdminAuthRequired 管理员操作鉴权中间件
// 当环境变量或 .env 中配置了 ADMIN_API_KEY 时，对写接口强制进行密钥核验
// 若未配置 ADMIN_API_KEY，放行并允许本地开发快速体验
func AdminAuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		expectedKey := strings.TrimSpace(config.AppConfig.AdminAPIKey)
		if expectedKey == "" {
			// 未设置密钥时放行（本地开发模式）
			c.Next()
			return
		}

		// 1. 尝试从 Authorization: Bearer <key> 头部获取
		authHeader := c.GetHeader("Authorization")
		token := ""
		if strings.HasPrefix(authHeader, "Bearer ") {
			token = strings.TrimSpace(strings.TrimPrefix(authHeader, "Bearer "))
		}

		// 2. 尝试从自定义 Header X-Admin-Key 获取
		if token == "" {
			token = strings.TrimSpace(c.GetHeader("X-Admin-Key"))
		}

		// 3. 尝试从 Query 参数获取 (备用)
		if token == "" {
			token = strings.TrimSpace(c.Query("admin_key"))
		}

		if token != expectedKey {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"message": "未授权的写操作：请在右上角「设置 -> 管理员钥匙」中输入正确的密钥",
				"data":    nil,
			})
			return
		}

		c.Next()
	}
}
