package router

import (
	"github.com/gin-gonic/gin"
	"github.com/goPractice/memoServer/controllers"
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
	r.GET("/v1", controllers.ReturnPageInfo)

	rr.router = r
	return &rr
}
