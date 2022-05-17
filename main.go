package main

import (
	"fmt"
	"gee"
	"net/http"
)

func main() {
	engine := gee.New()
	engine.GET("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})
	v1 := engine.Group("v1")
	{
		v1.GET("/hello", func(c *gee.Context) {
			c.String(http.StatusOK, "Hello %s, you are at %s", c.Query("name"), c.Path)
		})
		v1.GET("/hello/:name", func(c *gee.Context) {
			c.String(http.StatusOK, "Hello %s, it's %s", c.Param("name"), c.Path)
		})
	}

	v2 := engine.Group("v2")
	{
		v2.POST("/login", func(c *gee.Context) {
			c.JSON(http.StatusOK, gee.H{
				"username": c.PostForm("username"),
				"password": c.PostForm("password"),
			})
		})
		v2.GET("/assets/*filepath", func(c *gee.Context) {
			c.JSON(http.StatusOK, gee.H{"filepath": c.Param("filepath")})
		})
	}

	v3 := v2.Group("/nest")
	{
		v3.GET("/", func(c *gee.Context) {
			c.HTML(http.StatusOK, "<h1>This is a test for nested grouping</h1>")
		})
		v3.GET("/:name", func(c *gee.Context) {
			c.HTML(http.StatusOK, fmt.Sprintf("<div>%s</div>", c.Param("name")))
		})
	}
	engine.Run(":9999")

}
