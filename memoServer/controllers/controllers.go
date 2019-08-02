package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/goPractice/memoServer/model"
)

func ReturnPageInfo(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "3"))
	paginator := model.GetData(page, limit)
	ctx.JSON(200, paginator)
}
