package router

import (
	"github.com/gin-gonic/gin"
	"github.com/manyminds/api2go/routing"
	"net/http"
)

type ginRouter struct {
	router *gin.Engine
}

func (g ginRouter) Handler() http.Handler {
	return g.router
}

func (g ginRouter) Handle(protocol, route string, handler routing.HandlerFunc) {
	wrappedCallback := func(c *gin.Context) {
		params := map[string]string{}
		for _, p := range c.Params {
			params[p.Key] = p.Value
		}

		handler(c.Writer, c.Request, params, c.Keys)
	}

	g.router.Handle(protocol, route, wrappedCallback)
}

//Gin creates a new api2go router to use with the gin framework
func Gin(g *gin.Engine) routing.Routeable {
	return &ginRouter{router: g}
}