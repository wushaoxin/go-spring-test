package controllers

import (
    "github.com/gin-gonic/gin"
    "github.com/go-spring/spring-core/gs"
    "go-spring-test/go/controllers/health"
    "go-spring-test/go/controllers/home"
    "net/http"
)

func init() {

    gs.Object(new(home.Controller)).Init(func(c *home.Controller) {
        // 注册路由
        gs.GetMapping("/", c.Home)
        gs.GetMapping("/2", c.Home)
    })

    service := Controller{}.FileService
    println(service)

    gs.Object(new(health.Controller)).Init(func(c *health.Controller) {
        // 注册路由
        gs.GetMapping("/health", c.Health)
    })

}

func GetUserList(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "code": http.StatusOK,
        "msg":  "msg",
        "data": "data",
    })
}

type Controller struct {
    FileService gs.Context `autowire:""`
}
