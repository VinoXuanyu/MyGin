package main

import (
	"fmt"
	"jin"
	"math"
	"net/http"
	"strconv"
)

func main() {
	// Initialize a handler Engine
	engine := jin.New()

	// Register middlewares
	engine.Use(jin.Logger())   // Logs every request and it's time cost
	engine.Use(jin.Recovery()) // Avoid program shut down when Internal Server Error occurs

	// Register static route (only GET & POST supported by now)
	engine.GET("/admin/healthz", func(c *jin.Context) {
		c.String(http.StatusOK, "Welcome, this is Jin!") // Wraps string type response
	})

	engine.POST("/admin/healthz", func(c *jin.Context) {
		// Parse forms easily!
		val1 := c.PostForm("k1")
		val2 := c.PostForm("k2")
		// Wraps JSON type response
		c.JSON(http.StatusOK, jin.H{
			"This is v for k1": val1,
			"This is v for k2": val2,
		})
	})

	// Register dynamic route(':' to match single word and '*' to match all the resting parts)
	engine.GET("/api/hi/:name", func(c *jin.Context) {
		c.String(http.StatusOK, "Hi, This is %s", c.Param("name"))
	})

	// Route Grouping (groups share same prefix and also middlewares are segregated on group level)
	group1 := engine.Group("/api/group1")
	{
		group1.POST("/", func(c *jin.Context) {
			c.String(http.StatusOK, "Welcome, API Version 1!") //
		})
	}

	// Group nesting is also supported where prefixes are concatenated
	group2 := group1.Group("/nest")
	{
		group2.GET("/", func(c *jin.Context) {
			c.String(http.StatusOK, "This is a test for GET method in nested grouping")
		})
		group2.POST("/", func(c *jin.Context) {
			c.String(http.StatusOK, "This is a test for POST method in nested grouping")
		})
	}

	// Programme won't crash since we've registered Recovery middleware
	engine.GET("/admin/middleware-test", func(c *jin.Context) {
		var a []int
		fmt.Println("If you are seeing this something is wrong")
		fmt.Println(a[100])
	})

	// This is a test for performance of Golang
	engine.GET("/admin/stress-test", func(c *jin.Context) {
		temp := 0
		for i := 0; i < math.MaxInt32; i++ {
			temp += i
		}
		c.String(http.StatusOK, strconv.Itoa(temp))
	})

	engine.Run(":9999")

}
