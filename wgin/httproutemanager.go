/*
@Time : 2019/9/29 19:55 
@Author : liaoyizhong
@File : routemanager
@Software: GoLand
*/
package wgin

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type initRoute struct{
	Index int64
	Path string
	Name string
	Method string
	Ff func(*gin.Context)
}

var listInitRoute []initRoute

//初始化路由
func InitRouteRegister(){
	router := gin.Default()
	for _, v := range listInitRoute{
		switch v.Method {
		case "GET":
			router.GET(v.Path, v.Ff)
		case "POST":
			router.POST(v.Path, v.Ff)
		}
	}

	s := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()

}