package helper

import (
	"github.com/gin-gonic/gin"
	"github.com/holizDV/dts-go-chapter2-challenge3/app/domain/entity"
)

func HttpRespoonse(ctx *gin.Context, response entity.BaseHttpResponse) {
	ctx.JSON(response.Code, response)
}
