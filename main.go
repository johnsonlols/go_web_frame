package main

import "net/http"

func main() {
	// r := gin.Default()
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })
	// r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	e := New()
	e.Get("/", func(ctx *context) {
		rsp.Write([]byte{'/'})
	})

	e.Get("/panda", func(rsp http.ResponseWriter, req *http.Request) {
		rsp.Write([]byte{'p'})
	})

	e.Run(":9090")
}
