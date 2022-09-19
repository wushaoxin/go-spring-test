package health

import (
    "github.com/go-spring/spring-core/web"
)

type Controller struct {
}

func (c *Controller) Health(ctx web.Context) {
    ctx.JSON(web.SUCCESS.Data("Hello Go-Spring!"))
}
