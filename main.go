package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/inontz/kepler-go/router"
)

func main() {
	r := gin.Default()

	r.NoRoute(func(c *gin.Context) {
		sb := &strings.Builder{}
		sb.WriteString("routing err: no route, try this:\n")
		for _, v := range r.Routes() {
			sb.WriteString(fmt.Sprintf("%s %s\n", v.Method, v.Path))
		}
		c.String(http.StatusBadRequest, sb.String())
	})

	router.Routes(r)
	
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}