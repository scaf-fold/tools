package courier

import "github.com/gin-gonic/gin"

type Handler interface {
	Context(ctx *gin.Context)
}
