package jin

import (
	"fmt"
	"log"
	"runtime"
	"strings"
)

func traceback(message string) string {
	var pcs [32]uintptr
	n := runtime.Callers(3, pcs[:])

	var str strings.Builder
	str.WriteString(message + "\nTraceback:")
	for _, pc := range pcs[:n] {
		fn := runtime.FuncForPC(pc)
		file, line := fn.FileLine(pc)
		str.WriteString(fmt.Sprintf("\n\t%s:%d", file, line))
	}
	return str.String()
}

func Recover() HandlerFunc {
	return func(c *Context) {
		defer func() {
			if err := recover(); err != nil {
				message := fmt.Sprintf("%v", err)
				log.Printf("%s\n\n", traceback(message))
				c.Fail("Internal Server Error")
			}
		}()
		c.Next()
	}
}
