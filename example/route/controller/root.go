package controller

import (
	"github.com/scaf-fold/tools/courier"
	"github.com/scaf-fold/tools/example/route/controller/business"
)

// recommend using your service name
var root = courier.NewGroup("your base root")

var Root = &root

func init() {
	V1 := root.Group("/v1")
	{
		V1.Register(business.Root)
	}
}
