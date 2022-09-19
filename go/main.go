package main

import (
    "github.com/go-spring/spring-core/gs"
    // 初始化url
    _ "go-spring-test/go/controllers"
    // 这里必须引入web starter, 也可改为 "github.com/go-spring/starter-echo"
    _ "github.com/go-spring/starter-gin"
    "log"
)

func main() {
    // 启动应用程序, 同Java SpringBoot 的 SpringApplication.run(...)
    log.Fatal(gs.Run())
}
