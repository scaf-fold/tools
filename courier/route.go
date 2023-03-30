package courier

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type RouteGroup struct {
	root       string
	node       []interface{}
	nextRoot   []*RouteGroup
	middleWare []gin.HandlerFunc
}

func NewGroup(path string) *RouteGroup {
	return &RouteGroup{
		root: path,
	}
}

func RootRegister(g *gin.Engine, routes *RouteGroup) {
	if routes == nil {
		panic("routes is nil")
	}
	traversal(g.Group("/"), routes)
}

func (r *RouteGroup) Register(ctx interface{}) *RouteGroup {
	if c, ok := ctx.(*RouteGroup); ok {
		if r.nextRoot == nil {
			r.nextRoot = make([]*RouteGroup, 0)
		}
		r.nextRoot = append(r.nextRoot, c)
		return r
	}
	r.node = append(r.node, ctx)
	return r
}

func (r *RouteGroup) Group(path string, middleWare ...gin.HandlerFunc) *RouteGroup {
	child := NewGroup(path)
	for _, m := range middleWare {
		child.middleWare = append(child.middleWare, m)
	}
	if r.nextRoot == nil {
		r.nextRoot = make([]*RouteGroup, 0)
	}
	r.nextRoot = append(r.nextRoot, child)
	return child
}

func traversal(g *gin.RouterGroup, routes *RouteGroup) {
	if routes.root != "" {
		tmp := g.Group(routes.root)
		if routes.middleWare != nil && len(routes.middleWare) > 0 {
			tmp.Use(routes.middleWare...)
		}
		if routes.node != nil && len(routes.node) > 0 {
			for _, n := range routes.node {
				method := http.MethodOptions
				if mt, ok := n.(Method); ok {
					method = mt.String()
				}
				if handle, ok := n.(Context); ok {
					tmp.Handle(method, handle.Path(), handle.Context)
				}
			}
		}
		if routes.nextRoot != nil {
			for _, group := range routes.nextRoot {
				traversal(tmp, group)
			}
		}
	}
}
