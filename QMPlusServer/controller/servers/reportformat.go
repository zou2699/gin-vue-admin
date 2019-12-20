package servers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"gin-vue-admin/init/qmlog"
)

func ReportFormat(c *gin.Context, success bool, msg string, json gin.H) {
	// 开始时间
	start := time.Now()
	path := c.Request.URL.Path
	clientIP := c.ClientIP()
	method := c.Request.Method
	statusCode := c.Writer.Status()
	qmlog.QMLog.Infof("| %3d | %13v | %15s | %s  %s |%s|",
		statusCode,
		start,
		clientIP,
		method, path, gin.H{
			"success": success,
			"msg":     msg,
			"data":    json,
		},
	)
	c.JSON(http.StatusOK, gin.H{
		"success": success,
		"msg":     msg,
		"data":    json,
	})
}
