package courier

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var root = "/"

func RootRegister(g *gin.Engine, routes *RouteGroup) {
	if routes == nil {
		panic("routes is nil")
	}
	traversal(g.Group(root), routes)
}
func traversal(g *gin.RouterGroup, routes *RouteGroup) {
	if routes.root != "" {
		tmp := g.Group(routes.root)
		if routes.middleWare != nil && len(routes.middleWare) > 0 {
			tmp.Use(routes.middleWare...)
		}
		if routes.leafNode != nil && len(routes.leafNode) > 0 {
			for _, n := range routes.leafNode {
				method := http.MethodOptions
				if mt, ok := n.(Method); ok {
					method = mt.String()
				}
				if handle, ok := n.(Context); ok {
					tmp.Handle(method, handle.Path(), handle.Context)
				}
			}
		}
		if routes.subNode != nil {
			for _, group := range routes.subNode {
				traversal(tmp, group)
			}
		}
	}
}
