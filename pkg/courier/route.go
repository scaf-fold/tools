package courier

import (
	"github.com/gin-gonic/gin"
)

type RouteGroup struct {
	root       string
	leafNode   []interface{}
	subNode    []*RouteGroup
	middleWare []gin.HandlerFunc
}

func NewGroup(path string) *RouteGroup {
	return &RouteGroup{
		root: path,
	}
}

func (r *RouteGroup) Register(ctx interface{}) *RouteGroup {
	if c, ok := ctx.(*RouteGroup); ok {
		if r.subNode == nil {
			r.subNode = make([]*RouteGroup, 0)
		}
		r.subNode = append(r.subNode, c)
		return r
	}
	r.leafNode = append(r.leafNode, ctx)
	return r
}

func (r *RouteGroup) Group(path string, middleWare ...gin.HandlerFunc) *RouteGroup {
	child := NewGroup(path)
	for _, m := range middleWare {
		child.middleWare = append(child.middleWare, m)
	}
	if r.subNode == nil {
		r.subNode = make([]*RouteGroup, 0)
	}
	r.subNode = append(r.subNode, child)
	return child
}
