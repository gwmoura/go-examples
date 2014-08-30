package main

import "github.com/gin-gonic/gin"

func main() {
    r := gin.Default()
    r.GET("/ping", func(c *gin.Context) {
        c.String(200, "pong")
    })

    // However, this one will match /user/john and also /user/john/send
    r.GET("/user/:name/*action", func(c *gin.Context) {
        name := c.Params.ByName("name")
        action := c.Params.ByName("action")
        message := name + " is " + action
        c.String(200, message)
    })

    // Listen and server on 0.0.0.0:8080
    r.Run(":8080")
}