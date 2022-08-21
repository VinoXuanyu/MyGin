# Jin - A simplified Gin-like Web Framework 

## Design 
In most of the cases, when it comes to implement a web app, 
the first thing that comes in mind is what framework to choose. 
There are wide range of choices regardless of the programming language, 
like Flask, Django in Python; Spring, Mybatis in Java as well as Gin, Beego in Go.
But why don't we just build apps with standard libraries provided by the language SDK itself,
let's take a quick look on how to start a http service with `net/http` in Go,
which provides basic functionality of listening ports, mapping of static routes, and parsing HTTP requests,
but those are not enough for developing apps efficiently, we need some fancy things
like dynamic routing to capture certain params, middlewares to reduce duplicate code snippets,
and support for HTML templates 
```
func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe("localhost:8000", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello World!")
}
```

## Basics
Start an `engine` to register routes and corresponding handlers
```
func main() {
	r := jin.New()
	r.GET("/", func(c *jin.Context) {
		c.String(http.StatusOK, "Hiiiii")
	})
	r.Run(":9999")
}
```
## Features
### Context 
Provide uniformed way to parse requests and respond.
e.g. parse POST request and return JSON data 

- Without context
``` 
obj = map[string]interface{}{
    "jin": "good",
    "jinjin": "goodgood",
}
w.Header().Set("Content-Type", "application/json")
w.WriteHeader(http.StatusOK)
encoder := json.NewEncoder(w)
if err := encoder.Encode(obj); err != nil {
    http.Error(w, err.Error(), 500)
}
```
- With context 
``` 
c.JSON(http.StatusOK, jin.H{
    "jin": c.PostForm("jin"),
    "jinjin": c.PostForm("jinjin"),
})
```

### Dynamic routing
':' to match a single parameter, '*' for prefix match 
``` 
func main() {
    r := jin.New()
	r.GET("/hi/:name", func(c *jin.Context) {
		c.String(http.StatusOK, "hi %s", c.Param("name"))
	})
	r.Run(":9999")
```

### Routes grouping
Group routes with certain prefix to maintain well-formed API style.
``` 
func main() {
	r := jin.New()
	group1 := r.Group("/api/group1")
	{
		group1.GET("/", func(c *jin.Context) {
			c.String(http.StatusOK, "Hi. This is API Version 1.")
		})
	}
	v2 := r.Group("/api/v2")
	{
		v2.GET("/", func(c *jin.Context) {
			c.String(http.StatusOK, "Hi. This is API Version 2.")
		})
	}

	r.Run(":9999")
}
```

### Middlewares
Add middlewares to avoid duplicate codes like recording request metrics, logging, authing....
``` 
func main() {
	r := jin.New()
	r.Use(jin.Logger()) // global midlleware
	r.GET("/", func(c *jin.Context) {
		c.HTML(http.StatusOK, "<h1>Hello jin</h1>")
	})

```
