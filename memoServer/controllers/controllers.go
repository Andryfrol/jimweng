package controllers

import (
	"encoding/json"
	"log"
	"reflect"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jimweng/memoServer/model"
)

func ReturnPageInfo(ctx *gin.Context) {
	// parser querystring parameters
	page, _ := strconv.Atoi(ctx.DefaultQuery("offset", "1"))
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "3"))
	index := ctx.DefaultQuery("sort", "id")
	order := ctx.DefaultQuery("order", "asc")

	paginator := model.GetData(page, limit, index, order)
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
		title, description, category, ok := parseRequestBody(reqBody)
		if !ok {
			ctx.JSON(400, "{ status: 'Bad request'}")
			return
		} else {
			model.InsertData(title, description, category)
			ctx.JSON(201, "{ status: 'add data success'}")
			return
		}
	}
}

/* need to add log formater for better logger */
func UpdateData(ctx *gin.Context) {
	if reqBody, err := ctx.GetRawData(); err == nil {
		title, description, category, ok := parseRequestBody(reqBody)
		if !ok {
			ctx.JSON(400, "{ status: 'Bad request'}")
			return
		} else {
			model.UpdateData(title, description, category)
			ctx.JSON(200, "{ status: 'update data success'}")
			return
		}
	}
}

func parseRequestBody(reqBody []byte) (string, string, int, bool) {
	type f map[string]interface{}
	var postStructure f

	json.Unmarshal(reqBody, &postStructure)

	var title, description string
	var category int
	var ok bool
	for i, j := range postStructure {
		log.Printf("The value of i,j are %v_%v\n", i, j)
		switch i {
		case "title":
			title, ok = j.(string)
			log.Println(reflect.TypeOf(title))
			log.Printf("string %v\n", ok)

		case "description":
			description, ok = j.(string)
			log.Println(reflect.TypeOf(description))
			log.Printf("string %v\n", ok)

		case "category":
			category = int(j.(int))
			log.Println(reflect.TypeOf(category))
			log.Printf("int %v\n", ok)

		}
	}
	return title, description, category, ok
}
