package main

import (
	"fmt"
	"gee"
	"html/template"
	"log"
	"math"
	"net/http"
	"strconv"
)

// demo
func onlyForV2() gee.HandlerFunc {
	return func(c *gee.Context) {
		log.Printf("V2 middleware working !")
	}
}

func main() {
	engine := gee.Default()
	engine.LoadHTMLGlob("templates/*")
	engine.SetFuncMap(template.FuncMap{})
	engine.Static("/assets", "./static")

	engine.GET("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, "css.tmpl", nil)
	})

	engine.GET("/panic", func(c *gee.Context) {
		var a int
		slice := make([]int, 0)
		a = slice[100]
		fmt.Println(a)
	})

	engine.GET("stress", func(c *gee.Context) {
		temp := 0
		for i := 0; i < math.MaxInt32; i++ {
			temp += i
		}
		c.String(http.StatusOK, strconv.Itoa(temp))
	})

	v1 := engine.Group("/v1")
	{
		v1.GET("/hello", func(c *gee.Context) {
			c.String(http.StatusOK, "Hello %s, you are at %s", c.Query("name"), c.Path)
		})
		v1.GET("/hello/:name", func(c *gee.Context) {
			c.String(http.StatusOK, "Hello %s, it's %s", c.Param("name"), c.Path)
		})
	}

	v2 := engine.Group("/v2")
	v2.Use(onlyForV2())
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
			c.String(http.StatusOK, "This is a test for nested grouping")
		})
		v3.GET("/:name", func(c *gee.Context) {
			c.String(http.StatusOK, "Hi, %s", c.Param("name"))
		})
	}
	engine.Run(":9999")

}
