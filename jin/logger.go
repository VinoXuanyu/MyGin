package jin

import (
	"log"
	"time"
)

func Log() HandlerFunc {
	return func(c *Context) {
		t := time.Now()
		c.Next()
		log.Printf("[INFO][%d] %s in %v", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}
