package router

import (
	"github.com/gin-gonic/gin"
	"github.com/goPractice/memoServer/model"
)

type Router struct {
	router *gin.Engine
}

func (r *Router) Run() {
	r.router.Run()
}

type RouterImpl interface {
	Run()
}

func NewRouter() RouterImpl {
	var rr Router
	r := gin.Default()
	r.GET("/v1", model.ReturnPageInfo)

	rr.router = r
	return &rr
}
