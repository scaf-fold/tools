package business

import (
	"github.com/gin-gonic/gin"
	"github.com/scaf-fold/tools/courier"
	"net/http"
)

type BusinessExample struct {
	courier.MethodPut
	Id int64 `uri:"id"`
}

func (b *BusinessExample) Path() string {
	return "/:id"
}

func (b *BusinessExample) Context(ctx *gin.Context) {
	if err := ctx.ShouldBindUri(b); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
	}
	ctx.JSON(http.StatusNoContent, nil)
}
