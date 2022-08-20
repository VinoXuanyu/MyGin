package main

import (
	"fmt"
	"jin"
	"log"
	"math"
	"net/http"
	"strconv"
)

// demo
func onlyForV2() jin.HandlerFunc {
	return func(c *jin.Context) {
		log.Printf("V2 middleware working !")
	}
}

func main() {
	engine := jin.Default()
	engine.GET("/panic", func(c *jin.Context) {
		var a int
		slice := make([]int, 0)
		a = slice[100]
		fmt.Println(a)
	})

	engine.GET("/stress", func(c *jin.Context) {
		temp := 0
		for i := 0; i < math.MaxInt32; i++ {
			temp += i
		}
		c.String(http.StatusOK, strconv.Itoa(temp))
	})

	v1 := engine.Group("/v1")
	{
		v1.GET("/hello", func(c *jin.Context) {
			c.String(http.StatusOK, "Hello %s, you are at %s", c.Query("name"), c.Path)
		})
		v1.GET("/hello/:name", func(c *jin.Context) {
			c.String(http.StatusOK, "Hello %s, it's %s", c.Param("name"), c.Path)
		})
	}

	v2 := engine.Group("/v2")
	v2.Use(onlyForV2())
	{
		v2.POST("/login", func(c *jin.Context) {
			c.JSON(http.StatusOK, jin.H{
				"username": c.PostForm("username"),
				"password": c.PostForm("password"),
			})
		})
	}

	v3 := v2.Group("/nest")
	{
		v3.GET("/", func(c *jin.Context) {
			c.String(http.StatusOK, "This is a test for nested grouping")
		})
		v3.GET("/:name", func(c *jin.Context) {
			c.String(http.StatusOK, "Hi, %s", c.Param("name"))
		})
	}
	engine.Run(":9999")

}
