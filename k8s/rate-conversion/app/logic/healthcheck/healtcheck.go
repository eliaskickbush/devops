package healthcheck

import "github.com/gin-gonic/gin"

// Register ...
func Register(r *gin.Engine) {
	r.GET("/healthcheck", HealthCheck)
}
