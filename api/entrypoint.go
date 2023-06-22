package api

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/inontz/kepler-go/router"

	"github.com/gin-gonic/gin"
)

var (
	app *gin.Engine
)

// init gin app
func init() {
	app := gin.Default()

	// Handling routing errors
	app.NoRoute(func(c *gin.Context) {
		sb := &strings.Builder{}
		sb.WriteString("routing err: no route, try this:\n")
		for _, v := range app.Routes() {
			sb.WriteString(fmt.Sprintf("%s %s\n", v.Method, v.Path))
		}
		c.String(http.StatusBadRequest, sb.String())
	})
	
	router.Routes(app)

}

// entrypoint
func Handler(w http.ResponseWriter, r *http.Request) {
	app.ServeHTTP(w, r)
}
