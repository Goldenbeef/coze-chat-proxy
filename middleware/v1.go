package middleware

import (
	"coze-chat-proxy/config"
	"coze-chat-proxy/logger"
	"fmt"
	"github.com/gin-gonic/gin"
)

// V1Auth 验证v1 api 的token
func V1Auth(c *gin.Context) {
	authToken := c.Request.Header.Get("Authorization")
	// TODO: 验证token
	if authToken == "" {
		c.AbortWithStatusJSON(401, gin.H{
			"message": "You didn't provide an API key. You need to provide your API key in an Authorization header using Bearer auth (i.e. Authorization: Bearer YOUR_KEY).",
			"type":    "invalid_request_error",
			"param":   nil,
			"code":    nil,
		})
		return
	}
	if authToken != "Bearer "+config.CONFIG.AuthToken {
		c.AbortWithStatusJSON(401, gin.H{
			"message": "Incorrect API key provided: sk-4yNZz***************************************6mjw.",
			"type":    "invalid_request_error",
			"param":   nil,
			"code":    "invalid_api_key",
		})
		return
	}
	c.Next()
}

// V1Cors 跨域中间件
func V1Cors(c *gin.Context) {
	// 打印请求摘要 方法 url ip - user-agent 格式化输出
	infoStr := fmt.Sprint(c.Request.Method, " ", c.Request.URL.String(), " - ", c.ClientIP(), " - ", c.Request.Header.Get("User-Agent"))
	logger.Logger.Info(infoStr)
	// 允许跨域
	c.Header("Access-Control-Allow-Origin", "*")
	c.Next()
}

// V1Response 响应中间件
func V1Response(c *gin.Context) {
	c.Next()
	// 打印响应摘要 方法 url 状态码
	infoStr := fmt.Sprint(c.Request.Method, " ", c.Request.URL.String(), " - ", c.Writer.Status())
	logger.Logger.Info(infoStr)
}
