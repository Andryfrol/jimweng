package controllers

import (
	"encoding/json"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jimweng/memoServer/model"
)

func ReturnPageInfo(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "3"))
	paginator := model.GetData(page, limit)
	ctx.JSON(200, paginator)
	return
}

/* need to add log formater for better logger */
func DeleteSpecificValue(ctx *gin.Context) {
	id := ctx.DefaultQuery("id", "0")
	model.DeleteData(id)
	ctx.JSON(200, "{status: 'delete success'}")
	return
}

/* need to add log formater for better logger */
func PostData(ctx *gin.Context) {
	if reqBody, err := ctx.GetRawData(); err == nil {
		postStructure := model.MemoList{}
		json.Unmarshal(reqBody, &postStructure)
		model.InsertData(postStructure.Title, postStructure.Description, postStructure.Category)
	}
	ctx.JSON(201, "{ status: 'add data success}")
	return
}

/* need to add log formater for better logger */
func UpdateData(ctx *gin.Context) {
	if reqBody, err := ctx.GetRawData(); err == nil {
		postStructure := model.MemoList{}
		json.Unmarshal(reqBody, &postStructure)
		model.UpdateData(postStructure.Title, postStructure.Description, postStructure.Category)
	}
	ctx.JSON(200, "{ status: 'update data success}")
	return
}
